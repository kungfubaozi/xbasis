package main

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/afex/hystrix-go/hystrix"
	"golang.org/x/time/rate"
	"io/ioutil"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/errstate"
	"net"
	"net/http"
	"time"
)

var c *http.Client
var connectTimeout = 5 * time.Second
var readWriteTimeout = 100 * time.Millisecond
var configure = make(map[string]bool)
var limiters = make(map[string]*rate.Limiter)

func timeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func (r *request) route(req *http.Request) {

	limiter := limiters[r.serviceName]
	if limiter == nil {
		limiter = rate.NewLimiter(rate.Every(time.Duration(5)*time.Millisecond), 1)
		limiters[r.serviceName] = limiter
	}

	if err := limiter.Wait(context.Background()); err != nil {
		r.json(errstate.ErrOperateBusy)
		return
	}

	if c == nil {
		c = &http.Client{
			Transport: &http.Transport{
				Dial: timeoutDialer(connectTimeout, readWriteTimeout),
			},
		}
	}

	r.buildHeader(req)

	req.Header.Set("Content-Type", "application/json")

	r.services._log.WithFields(logrus.Fields{
		"service": r.serviceName,
		"routeTo": r.path,
	}).Info("request redirect")

	rt := time.Now().UnixNano()

	if !configure[r.serviceName] {
		hystrix.ConfigureCommand(r.serviceName, hystrix.CommandConfig{
			Timeout:                5000,
			MaxConcurrentRequests:  500,
			RequestVolumeThreshold: 3,
			ErrorPercentThreshold:  25,
			SleepWindow:            1000,
		})
	}

	hystrix.Do(r.serviceName, func() error {
		resp, err := c.Do(req)
		if err != nil {
			switch err {
			case http.ErrHandlerTimeout:
				r.json(errstate.ErrServiceRequestTimeout)
				return nil
			default:
				return err
			}
		}

		if r.dat != nil {
			from := encrypt.SHA1(r.rh.ip + r.rh.userAgent + r.rh.userDevice + r.rh.fromClientId)
			if r.dat.From != from {
				r.json(errstate.ErrDurationAccess)
				return nil
			}
			user := ""
			if r.auth {
				if len(r.userId) == 0 {
					r.json(errstate.ErrRequest)
					return nil
				}
				user = r.userId
			}
			if r.dat.FuncId != r.funcId || r.dat.ClientId != r.rh.fromClientId || (r.auth && user != r.dat.User) {
				r.json(errstate.ErrDurationAccess)
				return nil
			}

			s, err := r.services.accessibleService.DatReduce(r.ctx, r.dat)
			if err != nil {
				return err
			}
			if !s.State.Ok {
				r.json(s.State)
				return nil
			}
		}

		b, err := ioutil.ReadAll(resp.Body)
		r.c.Writer.Write(b)

		r.services._log.WithFields(logrus.Fields{
			"all":        fmt.Sprintf("%dms", (time.Now().UnixNano()-r.startAt)/1e6),
			"processing": fmt.Sprintf("%dms", (time.Now().UnixNano()-rt)/1e6),
		}).Info("all time consuming")

		return nil
	}, func(e error) error {
		r.services._log.WithFields(logrus.Fields{
			"err":     e.Error(),
			"service": r.serviceName,
		}).Error("request service failed")

		switch e {
		case hystrix.ErrCircuitOpen:
			r.json(errstate.ErrOperateBusy)
			break
		case hystrix.ErrMaxConcurrency:
			r.json(errstate.ErrOperateBusy)
			break
		case hystrix.ErrTimeout:
			r.json(errstate.ErrServiceRequestTimeout)
			break
		}

		r.json(errstate.ErrSystem)

		return nil
	})

}

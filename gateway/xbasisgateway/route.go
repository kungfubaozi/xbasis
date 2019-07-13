package xbasisgateway

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/afex/hystrix-go/hystrix"
	"golang.org/x/time/rate"
	"io/ioutil"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/commons/constants"
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
var limiter *rate.Limiter

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

	//if limiter == nil {
	//	limiter = rate.NewLimiter(rate.Every(time.Duration(5)*time.Millisecond), 1)
	//}
	//
	//if err := limiter.Wait(context.Background()); err != nil {
	//	r.json(errstate.ErrOperateBusy)
	//	return
	//}

	if c == nil {
		//c = &http.Client{
		//	Transport: &http.Transport{
		//		Dial: timeoutDialer(connectTimeout, readWriteTimeout),
		//	},
		//}
		c = http.DefaultClient
	}

	header := &analysisclient.LogHeaders{
		TraceId:     r.traceId,
		UserId:      r.userId,
		ServiceName: xbasisconstants.GatewayService,
		ModuleName:  "RouteToTarget",
	}

	r.buildHeader(req)

	req.Header.Set("Content-Type", "application/json")

	r.services._log.WithFields(logrus.Fields{
		"service": r.serviceName,
		"routeTo": r.path,
	}).Info("request redirect")

	r.services.log.Info(&analysisclient.LogContent{
		Headers: header,
		Action:  "StartRequestTargetAppApi",
		Message: fmt.Sprintf("%s Host: %s  Path: %s", r.requestMethod, r.path, r.requestPath),
		Fields: &analysisclient.LogFields{
			"function_id":  r.funcId,
			"route_to":     r.path,
			"path":         req.RequestURI,
			"service_name": r.serviceName,
		},
		Index: &analysisclient.LogIndex{
			Id:   r.logId,
			Name: r.logIndex,
			Fields: &analysisclient.LogFields{
				"route_to": r.path,
			},
		},
	})

	rt := time.Now().UnixNano()

	if !configure[r.serviceName] {
		hystrix.ConfigureCommand(r.serviceName, hystrix.CommandConfig{
			Timeout:                8000,
			MaxConcurrentRequests:  800,
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

		processing := (time.Now().UnixNano() - rt) / 1e6

		r.services._log.WithFields(logrus.Fields{
			"all":        fmt.Sprintf("%dms", (time.Now().UnixNano()-r.startAt)/1e6),
			"processing": fmt.Sprintf("%dms", processing),
		}).Info("all time consuming")

		all := (time.Now().UnixNano() - r.startAt) / 1e6

		r.services.log.Info(&analysisclient.LogContent{
			Headers:   header,
			Action:    "UserRequestApiFinished",
			StateCode: int64(resp.StatusCode),
			Message:   fmt.Sprintf("StatusCode: %d Timing: %dms", resp.StatusCode, processing),
			Fields: &analysisclient.LogFields{
				"function_id":   r.funcId,
				"function_name": r.funcName,
				"app_id":        r.funcAppId,
				"function_path": r.funcPath,
				"processing":    processing,
				"all":           all,
				"ok":            true,
				"verification":  all - processing,
			},
			Index: &analysisclient.LogIndex{
				Id:   r.logId,
				Name: r.logIndex,
				Fields: &analysisclient.LogFields{
					"processing":   processing,
					"all":          all,
					"verification": all - processing,
				},
			},
		})

		return nil
	}, func(e error) error {
		r.services._log.WithFields(logrus.Fields{
			"err":     e.Error(),
			"service": r.serviceName,
		}).Error("request service failed")

		all := (time.Now().UnixNano() - r.startAt) / 1e6
		processing := (time.Now().UnixNano() - rt) / 1e6

		r.services.log.Error(&analysisclient.LogContent{
			Headers:   header,
			Action:    "UserRequestApiFinished",
			StateCode: errstate.ErrRequest.Code,
			Message:   fmt.Sprintf("State: %s ERROR: %s", errstate.ErrRequest.Message, e.Error()),
			Fields: &analysisclient.LogFields{
				"function_id":   r.funcId,
				"function_name": r.funcName,
				"app_id":        r.funcAppId,
				"function_path": r.funcPath,
				"processing":    processing,
				"all":           all,
				"ok":            false,
				"verification":  all - processing,
			},
			Index: &analysisclient.LogIndex{
				Id:   r.logId,
				Name: r.logIndex,
				Fields: &analysisclient.LogFields{
					"processing":   processing,
					"all":          all,
					"verification": all - processing,
				},
			},
		})

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

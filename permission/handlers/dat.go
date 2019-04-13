package permissionhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/user/pb/ext"
	"math/rand"
	"strconv"
	"time"
)

type durationAccessService struct {
	pool    *redis.Pool
	session *mgo.Session
	*indexutils.Client
	configuration  *gs_commons_config.GosionConfiguration
	messageService gs_ext_service_user.MessageService
}

func (svc *durationAccessService) GetRepo() functionRepo {
	return functionRepo{Client: svc.Client, session: svc.session.Clone()}
}

type sendToUserFunc func(to, code string) *gs_commons_dto.State

//ip, NoneAuth
func (svc *durationAccessService) Datp(ctx context.Context, in *gs_service_permission.DurationAccessRequest, out *gs_service_permission.DurationAccessResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.To) == 0 || len(in.Path) == 0 {
			return nil
		}
		hkey := encrypt.SHA1(auth.IP + auth.ClientId)
		svc.dat(auth.IP, out, in.Path, in.To, auth.ClientId, auth.AppId, hkey, in.Code, svc.pool.Get(), func(to, code string) *gs_commons_dto.State {
			return svc.sendTo(ctx, to, code, 1)
		})
		return nil
	})
}

//user, AuthTypeOfToken
func (svc *durationAccessService) Datu(ctx context.Context, in *gs_service_permission.DurationAccessRequest, out *gs_service_permission.DurationAccessResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Path) == 0 {
			return nil
		}
		hkey := encrypt.SHA1(auth.User + auth.ClientId)
		svc.dat(auth.User, out, in.Path, auth.User, auth.ClientId, auth.AppId, hkey, in.Code, svc.pool.Get(), func(to, code string) *gs_commons_dto.State {
			return svc.sendTo(ctx, to, code, 2)
		})
		return nil
	})
}

func (svc *durationAccessService) sendTo(ctx context.Context, to, code string, t int64) *gs_commons_dto.State {
	s, err := svc.messageService.SendVerificationCode(ctx, &gs_ext_service_user.SendRequest{
		To:          to,
		Type:        t,
		Code:        code,
		MessageType: svc.configuration.SendVerificationCodeType,
	})
	if err != nil {
		return errstate.ErrSystem
	}
	return s.State
}

func (svc *durationAccessService) dat(user string, out *gs_service_permission.DurationAccessResponse, path, to, clientId,
	appId, hkey string, code int64, conn redis.Conn, toUser sendToUserFunc) {
	repo := svc.GetRepo()
	defer repo.Close()

	api, err := repo.FindApi(appId, path)
	if err != nil {
		out.State = errstate.ErrRequest
		return
	}

	write := func(dat *durationAccess) error {
		b, err := msgpack.Marshal(dat)
		if err != nil {
			return err
		}
		_, err = conn.Do("hset", hkey, api.Api, b)
		if err != nil {
			return err
		}
		return nil
	}

	addCode := func() {

		var ext int64

		if svc.configuration.DurationAccessTokenSendCodeToType == 1002 { //email
			var t int64
			t = 10 * 60
			if svc.configuration.EmailVerificationCodeExpiredTime > 0 {
				t = svc.configuration.EmailVerificationCodeExpiredTime
			}
			ext = t * 1e6
		} else if svc.configuration.DurationAccessTokenSendCodeToType == 1001 { //phone
			var t int64
			t = 10 * 60
			if svc.configuration.PhoneVerificationCodeExpiredTime > 0 {
				t = svc.configuration.PhoneVerificationCodeExpiredTime
			}
			ext = t * 1e6
		} else {
			ext = 10 * 60 * 1e6 //10min
		}

		dat := &durationAccess{
			Path:          path,
			ClientId:      clientId,
			User:          to,
			CreateAt:      time.Now().UnixNano(),
			CodeExpiredAt: ext + time.Now().UnixNano(),
			Code:          rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000),
		}

		if write(dat) != nil {
			out.State = errstate.ErrSystem
			return
		}

		out.State = toUser(to, strconv.FormatInt(dat.Code, 10))
	}

	b, err := redis.Bytes(conn.Do("hget", hkey, api.Api))
	if err != nil && err == redis.ErrNil {
		if code > 0 {
			out.State = errstate.ErrDurationAccess
			return
		}
		//code not send
		addCode()
		return
	}

	if err != nil {
		out.State = errstate.ErrRequest
		return
	}

	dat := &durationAccess{}
	err = msgpack.Unmarshal(b, dat)
	if err != nil {
		out.State = errstate.ErrRequest
		return
	}

	if dat.Code == -1 { // user already verify
		out.State = errstate.ErrVerificationCode
		return
	}

	if code < 1000000 && code > 100000 {
		//verify
		if dat.User != to && dat.ClientId != clientId && dat.Path != path && dat.Code != code {
			out.State = errstate.ErrDurationAccess
			return
		}
		if time.Now().UnixNano()-dat.CodeExpiredAt >= 0 {
			out.State = errstate.ErrDurationAccessExpired
			return
		}
		et, err := encrypt.AESEncrypt([]byte(user), []byte(svc.configuration.CurrencySecretKey))
		if err != nil {
			out.State = errstate.ErrSystem
			return
		}
		stat := string(et)
		dat.Stat = stat
		dat.Life = api.ValTokenLife

		//rewrite
		if write(dat) != nil {
			out.State = errstate.ErrSystem
			return
		}

		out.State = errstate.Success
		out.Dat = stat
	} else if code == 0 {
		//resend
		if time.Now().UnixNano()/1e9-dat.CreateAt/1e9 >= svc.configuration.DurationAccessTokenRetryTime {
			addCode()
			return
		}
		out.State = errstate.ErrDurationAccessTokenBusy
	} else {
		//err code
		out.State = errstate.ErrVerificationCode
	}
}

func NewDurationAccessService(pool *redis.Pool, session *mgo.Session, configuration *gs_commons_config.GosionConfiguration,
	messageService gs_ext_service_user.MessageService, client *indexutils.Client) gs_service_permission.DurationAccessHandler {
	return &durationAccessService{pool: pool, Client: client, session: session, configuration: configuration, messageService: messageService}
}

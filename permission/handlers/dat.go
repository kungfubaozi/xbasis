package permission_handlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
	"konekko.me/gosion/user/pb/nops"
	"math/rand"
	"strconv"
	"time"
)

type durationAccessService struct {
	pool          *redis.Pool
	session       *mgo.Session
	configuration *gs_commons_config.GosionConfiguration
	message       gs_nops_service_message.MessageService
}

func (svc *durationAccessService) GetRepo() permission_repositories.FunctionRepo {
	return permission_repositories.FunctionRepo{Conn: svc.pool.Get(), Session: svc.session.Clone()}
}

type sendToUserFunc func(to, code string) *gs_commons_dto.State

//ip, NoneAuth
func (svc *durationAccessService) Datp(ctx context.Context, in *gs_service_permission.DurationAccessRequest, out *gs_service_permission.DurationAccessResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.To) == 0 || len(in.Path) == 0 {
			return nil
		}
		hkey := gs_commons_encrypt.SHA1(auth.IP + auth.ClientId)
		svc.dat(out, in.Path, in.To, auth.ClientId, auth.AppId, hkey, in.Code, svc.pool.Get(), func(to, code string) *gs_commons_dto.State {
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
		hkey := gs_commons_encrypt.SHA1(auth.User + auth.ClientId)
		svc.dat(out, in.Path, auth.User, auth.ClientId, auth.AppId, hkey, in.Code, svc.pool.Get(), func(to, code string) *gs_commons_dto.State {
			return svc.sendTo(ctx, to, code, 2)
		})
		return nil
	})
}

func (svc *durationAccessService) sendTo(ctx context.Context, to, code string, t int64) *gs_commons_dto.State {
	s, err := svc.message.SendVerificationCode(ctx, &gs_nops_service_message.SendRequest{
		To:   to,
		Type: t,
		Code: code,
	})
	if err != nil {
		return errstate.ErrSystem
	}
	return s.State
}

func (svc *durationAccessService) dat(out *gs_service_permission.DurationAccessResponse, path, to, clientId,
	appId, hkey string, code int64, conn redis.Conn, toUser sendToUserFunc) {
	repo := svc.GetRepo()
	defer repo.Close()

	api, err := repo.FindApi(appId, path)
	if err != nil {
		out.State = errstate.ErrRequest
		return
	}

	addCode := func() {

		et, err := gs_commons_encrypt.AESEncrypt([]byte(to), []byte(svc.configuration.CurrencySecretKey))
		if err != nil {
			out.State = errstate.ErrSystem
			return
		}

		dat := &permission_repositories.DurationAccess{
			Path:     path,
			ClientId: clientId,
			Key:      string(et),
			User:     to,
			CreateAt: time.Now().UnixNano(),
			Life:     api.ValTokenLife,
			Code:     rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000),
		}
		out.State = toUser(to, strconv.FormatInt(dat.Code, 10))
	}

	b, err := redis.Bytes(conn.Do("hmget", hkey, api.ApiTag))
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

	dat := &permission_repositories.DurationAccess{}
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

	} else if code == 0 {
		//resend
		if time.Now().UnixNano()/1e9-dat.CreateAt/1e9 >= svc.configuration.DurationAccessTokenRetryTime {
			addCode()
			return
		}
		out.State = errstate.ErrDurationAccessTokenBusy
		return
	} else {
		//err code
		out.State = errstate.ErrVerificationCode
		return
	}
}

func NewDurationAccessService(pool *redis.Pool, session *mgo.Session, configuration *gs_commons_config.GosionConfiguration) gs_service_permission.DurationAccessHandler {
	return &durationAccessService{pool: pool, session: session, configuration: configuration}
}
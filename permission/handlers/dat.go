package permission_handers

import (
	"context"
	"encoding/base64"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/wrapper"
	"konekko.me/gosion/permission/pb"
	"konekko.me/gosion/permission/repositories"
	"strconv"
	"time"
)

type durationAccessService struct {
	pool          *redis.Pool
	session       *mgo.Session
	configuration *gs_commons_config.GosionConfiguration
}

func (svc *durationAccessService) GetRepo() permission_repositories.FunctionRepo {
	return permission_repositories.FunctionRepo{Conn: svc.pool.Get(), Session: svc.session.Clone()}
}

func (svc *durationAccessService) Try(ctx context.Context, in *gs_service_permission.DurationAccessRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {

		if len(in.Path) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		conn := svc.pool.Get()
		defer conn.Close()

		api, err := repo.FindApi(auth.AppId, in.Path)
		if err != nil {
			return nil
		}

		if len(api.AuthTypes) > 0 {
			for _, v := range api.AuthTypes {
				if v == gs_commons_constants.AuthTypeOfValcode {
					key := gs_commons_encrypt.SHA1(auth.User + auth.UserAgent + auth.UserDevice + auth.ClientId + auth.IP)
					m, err := redis.String(conn.Do("hmget", key, gs_commons_encrypt.SHA1(api.Api)))
					if err != nil || err == redis.ErrNil {
						err = nil
						m = ""
					}
					if err != nil {
						return nil
					}
					if len(m) > 0 {
						b, err := redis.Bytes(conn.Do("get", m))
						if err != nil || err == redis.ErrNil {
							//send
							return adat(auth.User, key, auth.ClientId, api, conn)
						}
						if err != nil {
							return nil
						}
						dat := &permission_repositories.DurationAccess{}
						err = msgpack.Unmarshal(b, dat)
						if err != nil {
							return nil
						}
						if dat.CreateAt-time.Now().UnixNano() < svc.configuration.DurationAccessTokenRetryTime*1e9 {
							return errstate.ErrDurationAccessTokenBusy
						}
						if dat.ExpiredAt-time.Now().UnixNano() <= 0 {
							//del old
							conn.Do("del", m)
							//send
							return adat(auth.User, key, auth.ClientId, api, conn)
						}
					}
					return adat(auth.User, key, auth.ClientId, api, conn)
				}
			}
		}
		return nil
	})
}

func adat(userId string, key string, clientId string, api *permission_repositories.Function, conn redis.Conn) *gs_commons_dto.State {
	stat := base64.StdEncoding.EncodeToString([]byte(gs_commons_encrypt.SHA1(
		strconv.FormatInt(time.Now().UnixNano(), 10) + clientId + api.Api)))
	dat := &permission_repositories.DurationAccess{
		UserId:    userId,
		CreateAt:  time.Now().UnixNano(),
		ExpiredAt: time.Now().UnixNano() + api.ValTokenLife,
		ClientId:  clientId,
		Path:      api.Api,
	}
	b, err := msgpack.Marshal(dat)
	if err != nil {
		return errstate.ErrRequest
	}
	_, err = conn.Do("hmset", key, stat)
	if err != nil {
		conn.Do("hdel", key)
		return errstate.ErrRequest
	}
	_, err = conn.Do("set", stat, b)
	if err != nil {
		return errstate.ErrRequest
	}
	return errstate.Success
}

func NewDurationAccessService(pool *redis.Pool, session *mgo.Session, configuration *gs_commons_config.GosionConfiguration) gs_service_permission.DurationAccessHandler {
	return &durationAccessService{pool: pool, session: session, configuration: configuration}
}

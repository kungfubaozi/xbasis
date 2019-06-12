package permissionhandlers

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/commons/regx"
	"konekko.me/gosion/commons/wrapper"
	external "konekko.me/gosion/permission/pb"
	"konekko.me/gosion/user/pb/inner"
	"math/rand"
	"time"
)

type durationAccessService struct {
	pool    *redis.Pool
	session *mgo.Session
	*indexutils.Client
	messageService gosionsvc_internal_user.MessageService
	log            analysisclient.LogClient
	zk             *zk.Conn
}

func (svc *durationAccessService) GetRepo() functionRepo {
	return functionRepo{Client: svc.Client, session: svc.session.Clone()}
}

func (svc *durationAccessService) Send(ctx context.Context, in *external.SendRequest, out *gs_commons_dto.Status) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Credential) > 0 {
			return errstate.ErrRequest
		}

		configuration := serviceconfiguration.Get()

		credential, es := svc.getCredential(in.Credential)
		if !es.Ok {
			return es
		}

		//
		////credential的有效时间为10s
		//if time.Now().UnixNano()-credential.Timestamp >= 10*1e6 {
		//	return errstate.ErrDurationAccessCredential
		//}

		to := in.To

		if !credential.FromAuth && len(in.To) <= 8 {
			return errstate.ErrDurationAccessTarget
		}

		if credential.FromAuth {
			to = auth.Token.UserId
		}

		hkey := encrypt.SHA256(to + credential.FuncId + auth.FromClientId)

		path := "gs.dat.lock/" + hkey
		var version int32
		invalid := false

		_, s, err := svc.zk.Get(path)
		if err != nil && err != zk.ErrInvalidPath {
			err = nil
			invalid = true
		}

		if err != nil {
			return errstate.ErrRequest
		}

		if err == nil {
			t := time.Now().Unix()

			if s != nil {
				t = s.Mtime
				version = s.Version
			}
			//limit
			if time.Now().Unix()-t < configuration.DurationAccessTokenRetryTime*1000 {
				return errstate.ErrDurationAccessTokenBusy
			}
		}

		repo := svc.GetRepo()
		defer repo.Close()

		api, err := repo.FindApiByPrimaryId(credential.FuncId)
		if err != nil {
			return errstate.ErrRequest
		}

		conn := svc.pool.Get()

		write := func(dat *durationAccess) error {
			b, err := msgpack.Marshal(dat)
			if err != nil {
				return err
			}
			_, err = conn.Do("hset", hkey, api.Id, b)
			if err != nil {
				return err
			}
			return nil
		}

		var ext int64

		if configuration.DurationAccessTokenSendCodeToType == 1002 { //email
			if !gs_commons_regx.Email(to) && !credential.FromAuth {
				return errstate.ErrFormatEmail
			}
			var t int64
			t = 10 * 60
			if configuration.EmailVerificationCodeExpiredTime > 0 {
				t = configuration.EmailVerificationCodeExpiredTime
			}
			ext = t * 1e6
		} else if configuration.DurationAccessTokenSendCodeToType == 1001 { //phone
			if !gs_commons_regx.Phone(to) && !credential.FromAuth {
				return errstate.ErrFormatPhone
			}
			var t int64
			t = 10 * 60
			if configuration.PhoneVerificationCodeExpiredTime > 0 {
				t = configuration.PhoneVerificationCodeExpiredTime
			}
			ext = t * 1e6
		} else {
			ext = 10 * 60 * 1e6 //10min
		}

		from := encrypt.SHA1(auth.IP + auth.UserAgent + auth.UserDevice + auth.FromClientId)

		dat := &durationAccess{
			Auth:          credential.FromAuth,
			FuncId:        api.Id,
			ClientId:      auth.FromClientId,
			From:          from,
			User:          to,
			CreateAt:      time.Now().UnixNano(),
			CodeExpiredAt: ext + time.Now().UnixNano(),
			Code:          rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000),
		}

		if write(dat) != nil {
			return errstate.ErrSystem
		}

		st, err := svc.messageService.SendVerificationCode(ctx, &gosionsvc_internal_user.SendRequest{
			To:          to,
			Auth:        credential.FromAuth,
			Code:        fmt.Sprintf("%d", dat.Code),
			MessageType: configuration.DurationAccessTokenSendCodeToType,
		})

		if err != nil {
			return errstate.ErrSystem
		}

		if !st.State.Ok {
			return st.State
		}

		data := []byte("")
		if invalid {
			_, err = svc.zk.Create(path, data, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
			if err != nil {
				return errstate.ErrRequest
			}
		}
		_, err = svc.zk.Set(path, data, version+1)

		if err != nil {
			return errstate.ErrRequest
		}

		return errstate.Success
	})
}

func (svc *durationAccessService) Verify(ctx context.Context, in *external.VerifyRequest, out *external.VerifyResponse) error {
	return gs_commons_wrapper.ContextToAuthorize(ctx, out, func(auth *gs_commons_wrapper.WrapperUser) *gs_commons_dto.State {
		if len(in.Credential) > 0 && len(in.To) <= 8 && (in.Code <= 1000000 && in.Code >= 100000) {
			return errstate.ErrRequest
		}
		configuration := serviceconfiguration.Get()

		credential, es := svc.getCredential(in.Credential)
		if !es.Ok {
			return es
		}

		hkey := encrypt.SHA256(in.To + credential.FuncId + auth.FromClientId)

		conn := svc.pool.Get()

		b, err := redis.Bytes(conn.Do("hget", hkey, credential.FuncId))
		if err != nil && err == redis.ErrNil {
			return errstate.ErrDurationAccessUnsentCode
		}
		if err != nil {
			return errstate.ErrRequest
		}

		da := &durationAccess{}
		err = msgpack.Unmarshal(b, &da)
		if err != nil {
			return errstate.ErrRequest
		}

		if in.To != da.User || in.Code != da.Code || auth.FromClientId != da.ClientId {
			return errstate.ErrDurationAccessCode
		}

		if da.CodeExpiredAt >= time.Now().UnixNano() {
			return errstate.ErrDurationAccessExpired
		}

		//generate token
		id := gs_commons_generator.NewIDG()
		tokenKey := id.Get()

		key, err := encrypt.AESEncrypt([]byte(tokenKey), []byte(configuration.CurrencySecretKey))
		if err != nil {
			return errstate.ErrSystem
		}

		_, err = conn.Do("del", hkey)
		if err != nil {
			return errstate.ErrSystem
		}

		repo := svc.GetRepo()
		defer repo.Close()

		api, err := repo.FindApiByPrimaryId(credential.FuncId)
		if err != nil {
			return errstate.ErrRequest
		}

		dat := &durationAccessToken{
			ClientId: da.ClientId,
			FuncId:   da.FuncId,
			User:     da.Key,
			Times:    0,
			Auth:     da.Auth,
			From:     da.From,
			MaxTimes: api.ValTokenTimes,
		}

		b, err = msgpack.Marshal(dat)
		if err != nil {
			return errstate.ErrRequest
		}

		_, err = conn.Do("hset", "dat."+da.FuncId, tokenKey, b)
		if err != nil {
			return errstate.ErrRequest
		}

		out.Dat = key

		return errstate.Success
	})
}

func (svc *durationAccessService) getCredential(credential string) (*durationAccessCredential, *gs_commons_dto.State) {
	configuration := serviceconfiguration.Get()

	var c *durationAccessCredential
	b, err := encrypt.AESDecrypt(credential, []byte(configuration.CurrencySecretKey))
	if err != nil {
		return nil, errstate.ErrSystem
	}

	err = msgpack.Unmarshal([]byte(b), &credential)
	if err != nil {
		return nil, errstate.ErrSystem
	}

	if c == nil || len(c.FuncId) == 0 {
		return nil, errstate.ErrSystem
	}
	return c, errstate.Success
}

func NewDurationAccessService(pool *redis.Pool, session *mgo.Session,
	messageService gosionsvc_internal_user.MessageService, client *indexutils.Client, log analysisclient.LogClient) external.DurationAccessHandler {
	return &durationAccessService{pool: pool, Client: client, session: session, messageService: messageService, log: log}
}

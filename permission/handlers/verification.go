package permissionhandlers

import (
	"context"
	"github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/metadata"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/application/pb/inner"
	"konekko.me/xbasis/authentication/pb/inner"
	xconfig "konekko.me/xbasis/commons/config"
	"konekko.me/xbasis/commons/config/call"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/encrypt"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/wrapper"
	inner "konekko.me/xbasis/permission/pb/inner"
	"konekko.me/xbasis/safety/pb"
)

type verificationService struct {
	pool                          *redis.Pool
	session                       *mgo.Session
	configuration                 *xconfig.GosionConfiguration
	innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService
	blacklistService              xbasissvc_external_safety.BlacklistService
	innerAuthService              xbasissvc_internal_authentication.AuthService
	*indexutils.Client
	log  analysisclient.LogClient
	_log *logrus.Logger
	id   generator.IDGenerator
}

type requestHeaders struct {
	authorization string
	userAgent     string
	userDevice    string
	ip            string
	refClientId   string //作为API share
	path          string
	dat           string
	fromClientId  string
}

func (svc *verificationService) GetRepo() *functionRepo {
	return &functionRepo{session: svc.session.Clone(), Client: svc.Client}
}

//application verify
//ip, userDevice blacklist verify
//api exists and authType verify
func (svc *verificationService) Check(ctx context.Context, in *inner.HasPermissionRequest, out *inner.HasPermissionResponse) error {
	//var wg sync.WaitGroup
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		md, ok := metadata.FromContext(ctx)
		svc.configuration = serviceconfiguration.Get()
		if len(svc.configuration.CurrencySecretKey) == 0 {
			return errstate.ErrAuthorization
		}

		if ok {

			traceId := md["transport-trace-id"]
			if len(traceId) > 10 {

				_, err := encrypt.AESDecrypt(traceId, []byte(svc.configuration.CurrencySecretKey))
				if err != nil {
					return nil
				}

				out.FromClient = auth.FromClientId
				out.RefClientId = auth.RefClientId
				out.TraceId = traceId
				out.Ip = auth.IP
				out.UserDevice = auth.UserDevice
				out.UserAgent = auth.UserAgent
				out.User = auth.User
				out.AppId = auth.AppId
				out.AppType = auth.AppType
				out.LogId = auth.LogId
				out.LogIndex = auth.LogIndex
				out.Platform = auth.Platform
				if auth.Access != nil {
					out.DatTo = auth.Access.To
					if auth.Access.Auth {
						out.DatAuth = 2
					}
				}
				if auth.Token != nil {
					out.Token = &inner.TokenInfo{
						UserId:   auth.Token.UserId,
						ClientId: auth.Token.ClientId,
						Platform: auth.Token.ClientPlatform,
						AppType:  auth.Token.AppType,
						AppId:    auth.Token.AppId,
						Relation: auth.Token.Relation,
					}
				}

				return errstate.SuccessTraceCheck
			}

		}

		return nil
	})
}

func NewVerificationService(pool *redis.Pool, session *mgo.Session,
	innerApplicationStatusService xbasissvc_internal_application.ApplicationStatusService, blacklistService xbasissvc_external_safety.BlacklistService,
	innerAuthService xbasissvc_internal_authentication.AuthService, client *indexutils.Client, logger analysisclient.LogClient) inner.VerificationHandler {
	return &verificationService{pool: pool, session: session, innerApplicationStatusService: innerApplicationStatusService, id: generator.NewIDG(),
		blacklistService: blacklistService, innerAuthService: innerAuthService, Client: client, log: logger, _log: logrus.New()}
}

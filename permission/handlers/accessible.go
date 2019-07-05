package permissionhandlers

import (
	"context"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/olivere/elastic"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	inner "konekko.me/xbasis/permission/pb/inner"
	"sync"
)

type accessibleService struct {
	*indexutils.Client
	log     analysisclient.LogClient
	session *mgo.Session
	pool    *redis.Pool
}

func (svc *accessibleService) GetDat(ctx context.Context, in *inner.GetDatRequest, out *inner.GetDatResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		conn := svc.pool.Get()

		b, err := redis.Bytes(conn.Do("hget", "dat."+in.FuncId, in.Key))
		if err != nil {
			return nil
		}

		dat := &DurationAccessToken{}

		err = msgpack.Unmarshal(b, &dat)
		if err != nil {
			return errstate.ErrSystem
		}

		out.Data = &inner.FunctionDat{
			MaxTimes: dat.MaxTimes,
			Times:    dat.Times,
			From:     dat.From,
			FuncId:   dat.FuncId,
			Auth:     dat.Auth,
			ClientId: dat.ClientId,
			User:     dat.User,
		}

		return errstate.Success
	})
}

func (svc *accessibleService) DatReduce(ctx context.Context, in *inner.FunctionDat, out *inner.DatReduceResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		conn := svc.pool.Get()

		dat := &DurationAccessToken{
			MaxTimes: in.MaxTimes,
			Times:    in.Times,
			From:     in.From,
			FuncId:   in.FuncId,
			Auth:     in.Auth,
			ClientId: in.ClientId,
			User:     in.User,
		}

		if dat.MaxTimes >= dat.Times {
			in.Times = in.Times + 1
			b, err := msgpack.Marshal(dat)
			if err != nil {
				return errstate.ErrSystem
			}
			_, err = conn.Do("hset", "dat."+dat.FuncId, in.Cv, b)
			if err != nil {
				return errstate.ErrSystem
			}
		} else {
			//delete access token
			conn.Do("del", "dat."+dat.FuncId, in.Cv)
			return errstate.ErrDurationAccessExpired
		}

		return errstate.Success
	})
}

func (svc *accessibleService) LookupApi(ctx context.Context, in *inner.LookupApiRequest, out *inner.LookupApiResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		repo := svc.GetFunctionRepo()
		defer repo.Close()

		f, err := repo.SimplifiedLookupApi(in.AppId, in.Path)
		if err != nil {
			return nil
		}

		out.AppId = f.AppId
		out.Path = f.Path
		out.Id = f.Id
		out.Name = f.Name
		out.ValTokenTimes = f.ValTokenTimes
		out.Share = f.Share
		out.GrantPlatforms = f.GrantPlatforms

		return errstate.Success
	})
}

func (svc *accessibleService) GetFunctionRepo() *functionRepo {
	return &functionRepo{Client: svc.Client}
}

func (svc *accessibleService) GetRepo() *bindingRepo {
	return &bindingRepo{session: svc.session.Clone(), Client: svc.Client}
}

func (svc *accessibleService) HasGrant(ctx context.Context, in *inner.HasGrantRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		r, err := repo.FindRelationUserById(in.UserId, in.AppId)
		if err != nil {
			return nil
		}

		if len(in.Role) > 10 {
			for _, v := range r.Roles {
				if v == in.Role {
					return errstate.ErrUserAlreadyBindRole
				}
			}
		}

		return errstate.Success
	})
}

func (svc *accessibleService) Check(ctx context.Context, in *inner.CheckRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		header := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ModuleName:  "Check",
			ServiceName: xbasisconstants.InternalPermission,
		}

		q := elastic.NewBoolQuery()
		q.Must(elastic.NewMatchPhraseQuery("userId", in.UserId), elastic.NewMatchPhraseQuery("functionId", in.FunctionId))
		e := svc.GetElasticClient().Search(getFunctionAuthorizeIndex(in.UserId)).Query(q)
		e.FetchSourceContext(elastic.NewFetchSourceContext(true).Include("recheck", "access"))
		v, err := e.Do(context.Background())
		if err != nil {
			return errstate.ErrRequest
		}
		a := &accessibale{}
		if v.Hits.TotalHits == 1 {
			err = json.Unmarshal(*v.Hits.Hits[0].Source, a)
			if err != nil {
				return errstate.ErrRequest
			}
		}

		if v.Hits.TotalHits > 1 {
			return errstate.ErrSystem
		}

		if v.Hits.TotalHits == 0 {
			a.Recheck = true
		}

		if !a.Access && !a.Recheck {
			return errstate.ErrUserPermission
		}

		//这里的目的就是对用户权限进行授权
		if a.Recheck {

			//先获取用户和功能的角色
			repo := svc.GetRepo()
			defer repo.Close()

			var wg sync.WaitGroup
			wg.Add(2)

			s := errstate.Success
			var f *functionRolesRelation
			var u *userRolesRelation

			resp := func(e *commons.State) {
				if s.Ok {
					s = e
				}
			}

			go func() {
				defer wg.Done()

				f1, err := repo.FindRelationFunctionById(in.FunctionId, in.AppId)
				if err != nil {
					resp(errstate.ErrSystem)
					return
				}

				f = f1

			}()

			go func() {
				defer wg.Done()

				u1, err := repo.FindRelationUserById(in.UserId, in.AppId)
				if err != nil {
					resp(errstate.ErrSystem)
					return
				}

				u = u1

			}()

			wg.Wait()

			if !s.Ok {
				return s
			}

			ok := false
			for _, v := range f.Roles {
				for _, v1 := range u.Roles {
					if v == v1 {
						ok = true
						break
					}
				}
			}

			if !ok {
				return errstate.ErrUserPermission
			}

			svc.log.Info(&analysisclient.LogContent{
				Headers: header,
				Action:  "FunctionAuthorize",
				Message: "sync function authorize",
				Fields: &analysisclient.LogFields{
					"appId":      in.AppId,
					"userId":     in.UserId,
					"functionId": in.FunctionId,
					"roles":      f.Roles,
					"recheck":    false,
					"access":     true,
				},
			})

		}

		return errstate.Success

	})
}

func NewAccessibleService(c *indexutils.Client, session *mgo.Session, log analysisclient.LogClient) inner.AccessibleHandler {
	return &accessibleService{Client: c, log: log, session: session}
}

package permissionhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/olivere/elastic"
	"gopkg.in/mgo.v2"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	"konekko.me/xbasis/commons/indexutils"
	"konekko.me/xbasis/commons/wrapper"
	external "konekko.me/xbasis/permission/pb"
)

type roleService struct {
	session *mgo.Session
	pool    *redis.Pool
	*indexutils.Client
	bindingService external.BindingService
}

func (svc *roleService) SearchUserRelations(ctx context.Context, in *external.SearchUserRelationsRequest, out *external.SearchUserRelationsResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *roleService) SearchFunctionRelations(ctx context.Context, in *external.SearchFunctionRelationsRequest, out *external.SearchFunctionRelationsResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		if len(in.RoleId) == 0 {
			return nil
		}

		//e := svc.GetElasticClient().Search(functionIndex)
		//query := elastic.NewBoolQuery()
		//
		//query.Must(elastic.NewMatchPhraseQuery("app_id", in.AppId))

		return nil
	})
}

func (svc *roleService) SearchRole(ctx context.Context, in *external.SearchRoleRequest, out *external.SearchRoleResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		if len(in.AppId) == 0 {
			return nil
		}

		e := svc.GetElasticClient().Search(roleIndex)

		query := elastic.NewBoolQuery()

		query.Must(elastic.NewMatchPhraseQuery("app_id", in.AppId))

		if len(in.Value) > 0 {
			q := elastic.NewQueryStringQuery("*" + in.Value + "*")
			if len(in.Key) > 0 {
				q.Field(in.Key)
			} else {
				q.Field("name")
			}
			query.Must(q)
		}

		v, err := e.Type("_doc").Query(query).From(int(in.Size * in.Page)).Size(int(in.Size)).Do(context.Background())
		if err != nil {
			return nil
		}

		var datas []*external.SimpleRoleInfo

		if v.Hits.TotalHits > 0 {
			for _, v := range v.Hits.Hits {
				i := &roleIndexModel{}
				err := json.Unmarshal(*v.Source, i)
				if err == nil {
					d := &external.SimpleRoleInfo{
						Name:      i.Name,
						Id:        i.Id,
						CreateAt:  i.CreateAt,
						Users:     i.RelationUsers,
						Functions: i.RelationFunctions,
					}
					datas = append(datas, d)
				}
			}
		}

		out.Data = datas

		return errstate.Success
	})
}

//修改为不分页
func (svc *roleService) GetAppRoles(ctx context.Context, in *external.GetAppRolesRequest, out *external.GetRoleResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		if len(in.AppId) == 0 {
			return nil
		}
		repo := svc.GetRepo()
		defer repo.Close()

		roles, err := repo.FindRolesByAppId(in.AppId, in.Page, in.Size)
		if err != nil {
			return nil
		}

		fmt.Println("data", len(roles))
		var rs []*external.SimpleRoleInfo
		for _, v := range roles {
			rs = append(rs, &external.SimpleRoleInfo{
				Id:       v.Id,
				Name:     v.Name,
				CreateAt: v.CreateAt,
			})
		}

		out.Data = rs
		return errstate.Success
	})
}

func (svc *roleService) GetRole(ctx context.Context, in *external.GetRoleRequest, out *external.GetRoleResponse) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		return nil
	})
}

func (svc *roleService) GetRepo() *roleRepo {
	return &roleRepo{session: svc.session.Clone(),
		id: generator.NewIDG(), conn: svc.pool.Get(), Client: svc.Client}
}

//add new role if not exists
func (svc *roleService) Add(ctx context.Context, in *external.RoleRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		repo := svc.GetRepo()
		defer repo.Close()

		_, err := repo.FindByName(in.Name, in.AppId)
		if err != nil && err == mgo.ErrNotFound {
			err = repo.Save(in.Name, auth.User, in.AppId)
			if err != nil {
				return nil
			}
			return errstate.Success
		}

		if err == nil {
			return errstate.ErrRoleAlreadyExists
		}

		return nil
	})
}

//remove role
//需要删除所有关联的角色对象包括(gs-user-roles-relation)
func (svc *roleService) Remove(ctx context.Context, in *external.RoleRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {

		return nil
	})
}

func (svc *roleService) Rename(ctx context.Context, in *external.RoleRequest, out *commons.Status) error {
	return xbasiswrapper.ContextToAuthorize(ctx, out, func(auth *xbasiswrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewRoleService(session *mgo.Session, pool *redis.Pool, bindingService external.BindingService, client *indexutils.Client) external.RoleHandler {
	return &roleService{session: session, pool: pool, bindingService: bindingService, Client: client}
}

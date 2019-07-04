package permissionhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	"konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	generator "konekko.me/xbasis/commons/generator"
	wrapper "konekko.me/xbasis/commons/wrapper"
	external "konekko.me/xbasis/permission/pb"
	"konekko.me/xbasis/user/pb/inner"
	"sync"
	"time"
)

type groupService struct {
	pool             *redis.Pool
	session          *mgo.Session
	innerUserService xbasissvc_internal_user.UserService
	log              analysisclient.LogClient
}

func (svc *groupService) GetGroupContentSize(ctx context.Context, in *external.GetGroupContentSizeRequest, out *external.GetGroupContentSizeResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.Id) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		var wg sync.WaitGroup
		wg.Add(2)
		var efg int64 = 0
		var efu int64 = 0
		s := errstate.Success
		resp := func(s1 *commons.State) {
			if s.Ok {
				s = s1
			}
		}

		go func() {
			defer wg.Done()
			c, err := repo.FindGroupItems(in.AppId, in.Id)
			if err != nil {
				resp(errstate.ErrRequest)
				return
			}
			efg = int64(len(c))
		}()

		go func() {
			defer wg.Done()
			c, err := repo.FindGroupUsers(in.AppId, in.Id)
			if err != nil {
				resp(errstate.ErrRequest)
				return
			}
			efu = int64(len(c))
		}()

		wg.Wait()

		if !s.Ok {
			return s
		}

		out.Users = efu
		out.Groups = efg

		return errstate.Success
	})

}

func (svc *groupService) GetGroupItems(ctx context.Context, in *external.GetGroupItemsRequest, out *external.GetGroupItemsResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.AppId) < 6 {
			return nil
		}

		header := &analysisclient.LogHeaders{
			TraceId:     auth.TraceId,
			ModuleName:  "GetGroupItems",
			ServiceName: xbasisconstants.PermissionService,
		}

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.Id) == 0 {
			in.Id = "-"
		}

		var groupItems []*external.GroupItem

		groups, err := repo.FindGroupItems(in.AppId, in.Id)
		if !mgoignore(err) {
			svc.log.Info(&analysisclient.LogContent{
				Headers: header,
				Action:  "FindGroupItemsMgoIgnore",
				Message: "mgo not found",
			})
			return nil
		}

		for _, v := range groups {
			groupItems = append(groupItems, &external.GroupItem{
				Id:   v.Id,
				Name: v.Name,
				User: false,
			})
		}

		if in.IncludeUser {

			users, err := repo.FindGroupUsers(in.AppId, in.Id)
			if !mgoignore(err) {
				return nil
			}

			var wg sync.WaitGroup
			if len(users) > 0 {
				s := errstate.Success
				resp := func(s1 *commons.State) {
					if s.Ok {
						s = s1
					}
				}

				getUserInfo := func(userId string) (string, bool) {
					s, err := svc.innerUserService.GetUserInfoById(ctx, &xbasissvc_internal_user.GetUserInfoByIdRequest{
						UserId: userId,
					})
					if err != nil {
						resp(errstate.ErrRequest)
						return "", false
					}
					if !s.State.Ok {
						resp(s.State)
						return "", false
					}
					resp(errstate.Success)
					return s.Username, true
				}

				if len(users) >= 2 {
					wg.Add(2)
					a := len(users) / 2

					go func() {
						defer wg.Done()
						a := users[:a]
						for _, v := range a {
							n, s := getUserInfo(v.UserId)
							if s {
								groupItems = append(groupItems, &external.GroupItem{
									Id:   v.UserId,
									User: true,
									Name: n,
								})
							}
						}
					}()

					go func() {
						defer wg.Done()
						a := users[:a]
						for _, v := range a {
							n, s := getUserInfo(v.UserId)
							if s {
								groupItems = append(groupItems, &external.GroupItem{
									Id:   v.UserId,
									User: true,
									Name: n,
								})
							}
						}
					}()
				} else {
					for _, v := range users {
						n, s := getUserInfo(v.UserId)
						if s {
							groupItems = append(groupItems, &external.GroupItem{
								Id:   v.UserId,
								User: true,
								Name: n,
							})
						}
					}
				}
				wg.Wait()

				if !s.Ok {
					return s
				}
			}
		}

		out.Data = groupItems

		return errstate.Success
	})
}

func (svc *groupService) GetGroupItemDetail(ctx context.Context, in *external.GetGroupItemDetailRequest, out *external.GetGroupItemDetailResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.AppId) < 8 {
			return nil
		}

		s, err := svc.innerUserService.GetUserInfoById(ctx, &xbasissvc_internal_user.GetUserInfoByIdRequest{
			UserId: in.Id,
		})

		if err != nil {
			return errstate.ErrRequest
		}

		out.Data = &external.DetailItem{
			Username: s.Username,
			RealName: s.RealName,
		}

		return errstate.Success
	})
}

func (svc *groupService) GetRepo() *groupRepo {
	return &groupRepo{session: svc.session.Clone(), id: generator.NewIDG()}
}

func (svc *groupService) Add(ctx context.Context, in *external.SimpleGroup, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		repo := svc.GetRepo()
		defer repo.Close()

		_, err := repo.FindByName(in.AppId, in.Name)
		if err != nil && err == mgo.ErrNotFound {

			if len(in.BindGroupId) == 0 {
				in.BindGroupId = "-"
			}

			id, err := repo.Save(in.AppId, auth.User, in.Name, in.BindGroupId)

			if err != nil {
				return errstate.ErrRequest
			}

			out.Content = id

			return errstate.Success
		}

		if err == nil {
			return errstate.ErrGroupAlreadyExists
		}

		return nil
	})
}

//重命名组
func (svc *groupService) Rename(ctx context.Context, in *external.SimpleGroup, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		return nil
	})
}

func (svc *groupService) AddUser(ctx context.Context, in *external.AddUserRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.AppId) < 5 && len(in.UserId) < 16 && len(in.GroupIds) == 0 {
			return nil
		}

		header := &analysisclient.LogHeaders{}

		repo := svc.GetRepo()
		defer repo.Close()

		ur, err := repo.FindUserById(in.UserId, in.AppId)
		if err != nil && err == mgo.ErrNotFound {
			ur = &userGroupsRelation{
				UserId:   in.UserId,
				AppId:    in.AppId,
				CreateAt: time.Now().UnixNano(),
			}
		}

		if err != nil {
			return nil
		}

		var groups []string
		for _, v := range in.GroupIds {
			ok := true
			for _, v1 := range ur.BindGroupId {
				if v1 == v {
					ok = false
					break
				}
			}
			if ok {
				groups = append(groups, v)
			}
		}

		if len(groups) > 0 {

			ur.BindGroupId = append(ur.BindGroupId, groups...)

			err = repo.SetGroupRelation(ur)
			if err != nil {
				return nil
			}

			for _, v := range groups {
				if v == "register" {
					svc.log.Info(&analysisclient.LogContent{
						Headers: header,
						Action:  "NewUser",
						Fields: &analysisclient.LogFields{
							"appId":  in.AppId,
							"userId": in.UserId,
						},
					})
				}
			}
		}

		return errstate.Success
	})
}

//移动用户
func (svc *groupService) MoveUser(ctx context.Context, in *external.SimpleUserNode, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.AppId) == 0 || len(in.UserId) == 0 {
			return nil
		}

		if len(in.GroupId) == 0 {
			in.GroupId = "-"
		}

		repo := svc.GetRepo()
		defer repo.Close()

		return nil
	})
}

//删除组
func (svc *groupService) Remove(ctx context.Context, in *external.SimpleGroup, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		return nil
	})
}

func NewGroupService(pool *redis.Pool, session *mgo.Session, innerUserService xbasissvc_internal_user.UserService) external.UserGroupHandler {
	return &groupService{pool: pool, session: session, innerUserService: innerUserService}
}

package permissionhandlers

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/xbasis/analysis/client"
	constants "konekko.me/xbasis/commons/constants"
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

func (svc *groupService) Search(context.Context, *external.SearchAppUserRequest, *external.SearchAppUserResponse) error {
	panic("implement me")
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
			ServiceName: constants.PermissionService,
		}

		repo := svc.GetRepo()
		defer repo.Close()

		if len(in.Id) == 0 {
			in.Id = constants.AppMainStructureGroup
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
				Id:       v.Id,
				Name:     v.Name,
				User:     false,
				GroupIds: []string{v.BindGroupId},
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

				getUserInfo := func(userId string, groups []string) {
					s, err := svc.innerUserService.GetUserInfoById(ctx, &xbasissvc_internal_user.GetUserInfoByIdRequest{
						UserId: userId,
					})
					if err != nil {
						resp(errstate.ErrRequest)
						return
					}
					if !s.State.Ok {
						resp(s.State)
						return
					}

					name := constants.GetStateString(s.UserState)

					groupItems = append(groupItems, &external.GroupItem{
						Id:       userId,
						User:     true,
						Name:     s.Username + name,
						GroupIds: groups,
					})
				}

				if len(users) >= 2 {
					wg.Add(2)
					a := len(users) / 2

					go func() {
						defer wg.Done()
						a := users[:a]
						for _, v := range a {
							getUserInfo(v.UserId, v.BindGroupId)
						}
					}()

					go func() {
						defer wg.Done()
						a := users[:a]
						for _, v := range a {
							getUserInfo(v.UserId, v.BindGroupId)
						}
					}()
				} else {
					for _, v := range users {
						getUserInfo(v.UserId, v.BindGroupId)
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
				in.BindGroupId = constants.AppMainStructureGroup
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

		header := &analysisclient.LogHeaders{
			ModuleName:  "AddUserToGroup",
			TraceId:     auth.TraceId,
			ServiceName: constants.PermissionService,
		}

		repo := svc.GetRepo()
		defer repo.Close()

		ur, err := repo.FindUserById(in.UserId, in.AppId)
		if err != nil && err == mgo.ErrNotFound {
			err = nil
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

			svc.log.Info(&analysisclient.LogContent{
				Headers: header,
				Action:  "AddUserToGroup",
				Fields: &analysisclient.LogFields{
					"userId":   in.UserId,
					"groupIds": in.GroupIds,
					"appId":    in.AppId,
					"operate":  auth.Token.UserId,
				},
				Index: &analysisclient.LogIndex{
					Id:   in.UserId,
					Name: "users",
					Fields: &analysisclient.LogFields{
						"app_" + in.AppId: true,
						"user_id":         in.UserId,
					},
				},
			})
		}

		return errstate.Success
	})
}

//移动用户或组
func (svc *groupService) Move(ctx context.Context, in *external.MoveRequest, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.AppId) == 0 || len(in.Id) == 0 || len(in.Groups) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		if !in.User {
			groupId := in.Groups[0]
			if groupId == in.Id {
				return nil
			}
			err := repo.MoveGroup(in.AppId, in.Id, groupId)
			if err != nil {
				return nil
			}
			return errstate.Success
		}

		err := repo.MoveUser(in.Id, in.AppId, in.Groups)
		if err == nil {
			return errstate.Success
		}

		return nil
	})
}

//删除组
func (svc *groupService) Remove(ctx context.Context, in *external.SimpleGroup, out *commons.Status) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {

		if len(in.Id) == 0 || len(in.AppId) == 0 {
			return nil
		}

		repo := svc.GetRepo()
		defer repo.Close()

		var wg sync.WaitGroup
		wg.Add(2)

		state := errstate.Success

		resp := func(e error) {
			if e == mgo.ErrNotFound {
				return
			}
			if state.Ok && e != nil {
				state = errstate.ErrRequest
			}
		}

		go func() {
			defer wg.Done()
			err := repo.DeleteUserGroupRelation(in.AppId, in.Id)
			resp(err)
		}()

		go func() {
			defer wg.Done()
			err := repo.DeleteGroupRelation(in.AppId, in.Id)
			resp(err)
		}()

		wg.Wait()

		return state
	})
}

func NewGroupService(pool *redis.Pool, session *mgo.Session, innerUserService xbasissvc_internal_user.UserService, log analysisclient.LogClient) external.UserGroupHandler {
	return &groupService{pool: pool, session: session, innerUserService: innerUserService, log: log}
}

package userhandlers

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	wrapper "konekko.me/xbasis/commons/wrapper"
	userpb "konekko.me/xbasis/user/pb"
	"strings"
)

type userService struct {
	client *indexutils.Client
}

func (svc *userService) Search(ctx context.Context, in *userpb.SearchRequest, out *userpb.SearchResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		if len(in.Value) == 0 || len(strings.TrimSpace(in.Value)) == 0 {
			return nil
		}
		q := elastic.NewQueryStringQuery(in.Value)
		if len(in.Key) > 0 {
			q.Field(in.Key)
		} else {
			q.Field("fields.username")
			q.Field("fields.real_name")
			q.Field("fields.account")
			q.Field("fields.phone")
			q.Field("fields.email")
		}

		v, err := svc.client.GetElasticClient().Search(typeUserIndex).Query(q).Type("_doc").From(int(in.Size * in.Page)).Size(int(in.Size)).Do(context.Background())
		if err != nil {
			return nil
		}

		var datas []*userpb.SimpleUserData

		if v.Hits.TotalHits > 0 {
			for _, v := range v.Hits.Hits {
				i := &userIndex{}
				err := json.Unmarshal(*v.Source, i)
				if err == nil {
					datas = append(datas, &userpb.SimpleUserData{
						UserId:   i.Fields.UserId,
						Username: i.Fields.Username,
						Icon:     i.Fields.Icon,
						State:    i.Fields.State,
						Invite:   i.Fields.Invite,
						Account:  i.Fields.Account,
						RealName: i.Fields.RealName,
					})
				}
			}
		}

		out.Data = datas

		return errstate.Success
	})
}

func NewUserService(client *indexutils.Client) userpb.UserHandler {
	return &userService{client: client}
}

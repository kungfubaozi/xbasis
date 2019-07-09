package userhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"konekko.me/xbasis/commons/constants"
	commons "konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/commons/indexutils"
	regx "konekko.me/xbasis/commons/regx"
	wrapper "konekko.me/xbasis/commons/wrapper"
	userpb "konekko.me/xbasis/user/pb"
)

type userService struct {
	client *indexutils.Client
}

func (svc *userService) Search(ctx context.Context, in *userpb.SearchRequest, out *userpb.SearchResponse) error {
	return wrapper.ContextToAuthorize(ctx, out, func(auth *wrapper.WrapperUser) *commons.State {
		//if len(in.AppId) == 0 && (len(in.Value) == 0 || len(strings.TrimSpace(in.Value)) == 0) {
		//	return nil
		//}

		if in.Key == "phone" && !regx.Phone(in.Value) {
			return errstate.ErrFormatPhone
		}
		if in.Key == "email" && !regx.Email(in.Value) {
			return errstate.ErrFormatEmail
		}
		v1 := in.Value
		if in.Card {
			v1 = fmt.Sprintf("*%s*", v1)
		}

		e := svc.client.GetElasticClient().Search(typeUserIndex)

		query := elastic.NewBoolQuery()

		if len(v1) > 0 {
			q := elastic.NewQueryStringQuery(v1)
			if len(in.Key) > 0 {
				q.Field(in.Key)
			} else {
				q.Field("username")
				q.Field("real_name")
				q.Field("account")
				q.Field("phone")
				q.Field("email")
			}
			query.Must(q)
		}
		if len(in.AppId) > 0 {
			query.Must(elastic.NewMatchPhraseQuery("app_"+in.AppId, true))
		}
		if in.Invite {
			query.Must(elastic.NewMatchPhraseQuery("from_invite", true))
		}

		v, err := e.Type("_doc").Query(query).From(int(in.Size * in.Page)).Size(int(in.Size)).Do(context.Background())
		if err != nil {
			return nil
		}

		var datas []*userpb.SimpleUserData

		if v.Hits.TotalHits > 0 {
			for _, v := range v.Hits.Hits {
				i := &userIndexFields{}
				err := json.Unmarshal(*v.Source, i)
				if err == nil {
					d := &userpb.SimpleUserData{
						UserId:     i.UserId,
						Username:   i.Username,
						Icon:       i.Icon,
						State:      int64(i.State),
						Invite:     i.Invite,
						Account:    i.Account,
						RealName:   i.RealName,
						FromInvite: i.FromInvite,
					}
					d.Label = xbasisconstants.GetStateString(d.State)
					datas = append(datas, d)
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

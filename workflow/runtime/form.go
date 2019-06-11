package runtime

import (
	"context"
	"github.com/garyburd/redigo/redis"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/analysis/client"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/flowerr"
	"konekko.me/gosion/workflow/models"
	"konekko.me/gosion/workflow/types"
	"time"
)

type form struct {
	session   *mgo.Session
	pool      *redis.Pool
	client    *indexutils.Client
	log       analysisclient.LogClient
	id        gs_commons_generator.IDGenerator
	secretKey string
}

func (f *form) Submit(ctx context.Context, instanceId, nodeId, formId string, encryption bool, value map[string]interface{}) *flowerr.Error {
	user := getWrapperUser(ctx)

	b, err := msgpack.Marshal(value)
	if err != nil {
		return flowerr.FromError(err)
	}

	if encryption {
		if len(f.secretKey) == 0 {
			return flowerr.ErrSecretKey
		}
		s, err := encrypt.AESEncrypt(b, []byte(f.secretKey))
		if err != nil {
			return flowerr.FromError(err)
		}
		b = []byte(s)
	}

	s := &models.SubmitForm{
		Info: &models.Info{
			Id:           f.id.String(),
			CreateAt:     time.Now().UnixNano(),
			CreateUserId: user.Token.UserId,
		},
		Encryption: encryption,
		FormId:     formId,
		NodeId:     nodeId,
		Data:       string(b),
		InstanceId: instanceId,
	}

	ok, err := f.client.AddData(types.IndexSubmitForm, s)
	if err != nil {
		return flowerr.FromError(err)
	}
	if len(ok) > 0 {
		err := f.session.DB(types.DBFlow).C(types.GetSubmitFormCollection(instanceId, nodeId)).Insert(s)
		if err != nil {
			return flowerr.FromError(err)
		}
		return nil
	}
	return flowerr.ErrUnknow
}

func (f *form) FindById(id string) (*models.TypeForm, *flowerr.Error) {
	panic("implement me")
}

func (f *form) LoadNodeDataFromStore(ctx context.Context, instanceId, nodeId string) (map[string]interface{}, *flowerr.Error) {
	var s *models.SubmitForm
	ok, err := f.client.QueryFirst(types.IndexSubmitForm, map[string]interface{}{"instance_id": instanceId, "node_id": nodeId}, &s)
	if err != nil {
		return nil, flowerr.FromError(err)
	}
	if ok {
		m := make(map[string]interface{})
		if s.Encryption {
			b, err := encrypt.AESDecrypt(s.Data, []byte(f.secretKey))
			if err != nil {
				return nil, flowerr.FromError(err)
			}
			s.Data = b
		}
		err = msgpack.Unmarshal([]byte(s.Data), m)
		if err != nil {
			return nil, flowerr.FromError(err)
		}
		return m, nil
	}
	return nil, nil
}

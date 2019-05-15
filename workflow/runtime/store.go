package runtime

import (
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/flowerr"
	"time"
)

type store struct {
	client *indexutils.Client
	log    *gslogrus.Logger
}

func (s *store) Clear(keys map[string]interface{}) *flowerr.Error {
	ok, err := s.client.Delete(storeIndex, keys)
	if err != nil {
		return flowerr.FromError(err)
	}
	if ok {
		return nil
	}
	return flowerr.ErrUnknow
}

var storeIndex = "gs-flow-store"

func (s *store) Get(keys map[string]interface{}) (int64, *flowerr.Error) {
	var v map[string]interface{}
	ok, err := s.client.QueryFirst(storeIndex, keys, &v, "status")
	if err != nil {
		return 0, flowerr.FromError(err)
	}
	if ok {
		i := v["status"].(int64)
		return i, nil
	}
	return 0, flowerr.ErrNil
}

func (s *store) Set(status int64, keys map[string]interface{}) (bool, *flowerr.Error) {
	keys["status"] = status
	keys["time"] = time.Now().UnixNano()
	ok, err := s.client.AddData(storeIndex, keys)
	if err != nil {
		return false, flowerr.FromError(err)
	}
	if len(ok) > 0 {
		return true, nil
	}
	return false, flowerr.ErrRequest
}

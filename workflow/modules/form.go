package modules

import (
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"konekko.me/gosion/commons/gslogrus"
	"konekko.me/gosion/commons/indexutils"
	"konekko.me/gosion/workflow/models"
)

type IForm interface {
	FindById(id string) (*models.TypeForm, error)
}

type form struct {
	session *mgo.Session
	pool    *redis.Pool
	client  *indexutils.Client
	log     *gslogrus.Logger
}

func (f *form) FindById(id string) (*models.TypeForm, error) {
	panic("implement me")
}

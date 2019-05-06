package models

type Info struct {
	Id           string `bson:"id" json:"id"`
	Key          string `bson:"key" json:"key"`
	Name         string `bson:"name" json:"name"`
	Desc         string `bson:"desc" json:"desc"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
}

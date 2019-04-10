package userhandlers

import (
	"fmt"
)

const (
	dbName = "gs_users"

	userCollection = "users"

	userInfoCollection = "user_info"

	userOAuthCollection = "user_oauth"
)

const (
	emailIndexType   = 2
	phoneIndexType   = 3
	accountIndexType = 4
)

type userInfo struct {
	Id         string `bson:"_id"`
	CreateAt   int64  `bson:"create_at"` //register time
	Password   string `bson:"password"`
	PIN        int64  `bson:"pin"` //number
	ModifyAt   int64  `bson:"modify_at"`
	RegisterAt string `bson:"register_at"` //register at clientId
	Account    string `bson:"account"`
	Email      string `bson:"email"`
	Phone      string `bson:"phone"`
}

type userOAuth struct {
	OpenId   string
	Type     int64
	CreateAt int64
}

type userPersonInfo struct {
	UserId   string `bson:"user_id"`
	Icon     string `bson:"icon"`
	FromCity string `bson:"from_city"`
	Birthday string `bson:"birthday"`
	Age      int32  `bson:"age"`
	Sex      int32  `bson:"sex"`
	Username string `bson:"username"`
	RealName string `bson:"real_name"`
	Desc     string `bson:"desc"`
	ModifyAt int64  `bson:"modify_at"`
}

type userIndex struct {
	Content  string `gorm:"index"`
	TargetId string
	Type     int
	Code     int
}

func (u *userIndex) TableName() string {
	return fmt.Sprintf("user_index_%d_%d", u.Type, u.Code)
}

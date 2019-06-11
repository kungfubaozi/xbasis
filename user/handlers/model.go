package userhandlers

import "konekko.me/gosion/commons/encrypt"

const (
	dbName = "gs_users"

	userCollection = "users"

	userInfoCollection = "user_info"

	userOAuthCollection = "user_oauth"
)

const (
	typeUserIndex = "gs-users"
)

const (
	Boy  = 3 << 2
	Girl = 3 << 4
)

type userOAuth struct {
	OpenId   string
	Type     int64
	CreateAt int64
}

type userPersonInfo struct {
	Icon     string `bson:"icon"`
	FromCity string `bson:"from_city"`
	Birthday int64  `bson:"birthday"`
	Age      int32  `bson:"age"`
	Sex      int32  `bson:"sex"`
	Username string `bson:"username"`
	RealName string `bson:"real_name"`
	Desc     string `bson:"desc"`
	ModifyAt int64  `bson:"modify_at"`
}

type userModel struct {
	Id         string         `bson:"_id"`
	SID        string         `bson:"sid"`
	CreateAt   int64          `bson:"create_at"`
	Account    string         `bson:"account"`
	Phone      string         `bson:"phone"`
	Email      string         `bson:"email"`
	Password   string         `bson:"password"`
	RegisterAt string         `bson:"register_at"`
	OAuth      []userOAuth    `bson:"o_auth"`
	Info       userPersonInfo `bson:"info"`
}

type userModelIndex struct {
	UserId     string `json:"user_id"`
	Account    string `json:"account"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	RealName   string `json:"real_name"`
	Sex        int32  `json:"sex"`
	Birthday   int64  `json:"birthday"`
	RegisterAt string `json:"register_at"`
	Age        int32  `json:"age"`
}

func (um *userModel) Index() *userModelIndex {
	index := &userModelIndex{UserId: um.Id}
	if len(um.Email) > 6 {
		index.Email = encrypt.SHA1(um.Email)
	}
	if len(um.Phone) > 6 {
		index.Phone = encrypt.SHA1(um.Phone)
	}
	if len(um.Account) > 4 {
		index.Account = encrypt.SHA1(um.Account)
	}
	if len(um.RegisterAt) > 10 {
		index.RegisterAt = um.RegisterAt
	}
	if len(um.Info.Username) > 2 {
		index.Username = um.Info.Username
	}
	if len(um.Info.RealName) > 2 {
		index.RealName = um.Info.RealName
	}
	if um.Info.Sex == Boy || um.Info.Sex == Girl {
		index.Sex = um.Info.Sex
	}
	if um.Info.Age > 0 {
		index.Age = um.Info.Age
	}
	if um.Info.Birthday > 0 {
		index.Birthday = um.Info.Birthday
	}
	return index
}

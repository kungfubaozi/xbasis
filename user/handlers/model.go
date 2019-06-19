package userhandlers

const (
	dbName = "gs_users"

	userCollection = "users"

	userInfoCollection = "user_info"

	userOAuthCollection = "user_oauth"

	inviteCollection = "invites"
)

const (
	typeUserIndex = "gs-users"
)

const (
	Boy  = 3 << 2
	Girl = 3 << 4
)

type inviteModel struct {
	Phone        string        `bson:"phone"`
	Email        string        `bson:"email"`
	CreateAt     int64         `bson:"create_at"`
	CreateUserId string        `bson:"create_user_id"`
	Username     string        `bson:"username"`
	RealName     string        `bson:"real_name"`
	Items        []*inviteItem `bson:"items"`
	Type         int64         `bson:"type"` //邀请类型
	UserId       string        `bson:"user_id"`
}

type inviteItem struct {
	AppId       string   `bson:"app_id"`
	BingGroupId string   `bson:"bing_group_id"`
	Roles       []string `bson:"roles"`
}

type userOAuth struct {
	OpenId   string `bson:"open_id" json:"open_id"`
	UnionId  string `bson:"union_id" json:"union_id"`
	Type     int64  `bson:"type" json:"type"` //facebook, wechat, qq, dingding, google
	UserId   string `bson:"user_id" json:"user_id"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}

type userInfo struct {
	UserId   string `bson:"user_id" json:"user_id"`
	Icon     string `bson:"icon" json:"icon"`
	FromCity string `bson:"from_city" json:"from_city"`
	Birthday int64  `bson:"birthday" json:"birthday"`
	Age      int32  `bson:"age" json:"age"`
	Sex      int32  `bson:"sex" json:"sex"`
	Username string `bson:"username" json:"username"`
	RealName string `bson:"real_name" json:"real_name"`
	Desc     string `bson:"desc" json:"desc"`
	ModifyAt int64  `bson:"modify_at" json:"modify_at"`
	CreateAt int64  `bson:"create_at" json:"create_at"`
}

type userModel struct {
	Id         string `bson:"_id"`
	CreateAt   int64  `bson:"create_at"`
	Account    string `bson:"account"`
	Phone      string `bson:"phone"`
	Email      string `bson:"email"`
	Password   string `bson:"password"`
	RegisterAt string `bson:"register_at"`
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
	index.Email = um.Email
	index.Phone = um.Phone
	index.Account = um.Account
	index.RegisterAt = um.RegisterAt

	return index
}

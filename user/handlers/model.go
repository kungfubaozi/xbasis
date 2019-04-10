package userhandlers

const (
	dbName = "gs_users"

	userCollection = "users"

	userInfoCollection = "user_info"

	userOAuthCollection = "user_oauth"
)

const (
	emailIndexType   = 2 << 3
	phoneIndexType   = 3 << 9
	accountIndexType = 4 << 5
)

const (
	typeUserIndex = "gs_user_index"
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
	Content  string `json:"content"`
	TargetId string `json:"target_id"`
}

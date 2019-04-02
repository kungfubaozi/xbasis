package userhandlers

type userInfo struct {
	Id         string `bson:"_id"`
	CreateAt   int64  `bson:"create_at"` //register time
	CardId     string `bson:"card_id"`
	Account    string `bson:"account"`
	Password   string `bson:"password"`
	PIN        int64  `bson:"pin"` //number
	ModifyAt   int64  `bson:"modify_at"`
	RegisterAt string `bson:"register_at"` //register at clientId
}

//The number of types should not be greater than 1
type userContractInfo struct {
	UserId   string `bson:"user_id"`
	Contract string `bson:"contract"`
	Type     int64  `bson:"type"`
	CreateAt int64  `bson:"create_at"`
	ModifyAt int64  `bson:"modify_at"`
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

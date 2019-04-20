package flowform

type TextForm struct {
	Id           string `bson:"_id" json:"id"`
	Name         string `bson:"name" json:"name"`
	Desc         string `bson:"desc" json:"desc"`
	Field        string `json:"field" bson:"field"`
	Type         int32  `json:"type" bson:"type"`
	CreateAt     int64  `json:"create_at" bson:"create_at"`
	CreateUserId string `json:"create_user_id" bson:"create_user_id"`
	Require      bool   `bson:"require" json:"require"`
}

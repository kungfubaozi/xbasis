package permission_repositories

type UserRole struct {
	Id            string          `bson:"_id"`
	CreateUserId  string          `bson:"create_user_id"`
	CreateAt      int64           `bson:"create_at"`
	Name          string          `bson:"name"`
	LinkRoles     []*LinkRole     `bson:"link_roles"`
	LinkAppGroups []*LinkAppGroup `bson:"link_app_groups"`
}

type UserGroup struct {
	Id            string          `bson:"_id"`
	CreateUserId  string          `bson:"create_user_id"`
	CreateAt      int64           `bson:"create_at"`
	Name          string          `bson:"name"`
	Type          int64           `bson:"type"`
	LinkAppGroups []*LinkAppGroup `bson:"link_app_groups"`
}

type LinkRole struct {
	AppId string   `bson:"app_id"`
	Roles []string `bson:"roles"`
}

type LinkAppGroup struct {
	AppId       string `bson:"app_id"`
	BindGroupId string `bson:"bind_group_id"`
}

type FunctionGroup struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	CreateUserId string `bson:"create_user_id"`
	CreateAt     int64  `bson:"create_at"`
	BindGroupId  string `bson:"bind_group_id"`
	Type         int64  `bson:"type"`
	AppId        string `bson:"app_id"`
}

type Function struct {
	Id           string  `bson:"_id"`
	Name         string  `bson:"name"`
	Api          string  `bson:"api"`
	ApiTag       string  `bson:"api_tag"`
	Type         int64   `bson:"type"`
	CreateUserId string  `bson:"create_user_id"`
	CreateAt     int64   `bson:"create_at"`
	BindGroupId  string  `bson:"bind_group_id"`
	AppId        string  `bson:"app_id"`
	AuthTypes    []int64 `bson:"auth_types"`
	//authType container AuthTypeOfValcode. valTokenLife is access this function token expired time
	ValTokenLife    int64   `bson:"val_token_life"` //def: 0 second (your value must >= 60s)
	NoGrantPlatform []int64 `bson:"no_grant_platform"`
}

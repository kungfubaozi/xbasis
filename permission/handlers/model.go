package permissionhandlers

type cacheStructure struct {
	UserStructureId     string
	FunctionStructureId string
}

type durationAccess struct {
	User          string
	Path          string
	Life          int64
	ClientId      string
	CreateAt      int64
	Stat          string
	Code          int64
	CodeExpiredAt int64
	Key           string
}

type structure struct {
	Id           string `bson:"_id"`
	CreateAt     int64  `bson:"create_at"`
	CreateUserId string `bson:"create_user_id"`
	AppId        string `bson:"app_id"`
	Opening      bool   `bson:"opening"`
	Name         string `bson:"name"`
	Type         int64  `bson:"type"` //user or function structure
}

type userOrientate struct {
	Id                  string                `bson:"_id"`
	CreateAt            int64                 `bson:"create_at"`
	UserId              string                `bson:"user_id"`
	LinkStructureRoles  []*linkStructureRole  `bson:"link_structure_roles"`
	LinkStructureGroups []*linkStructureGroup `bson:"link_structure_groups"`
}

type userGroup struct {
	Id                  string                `bson:"_id"`
	CreateUserId        string                `bson:"create_user_id"`
	CreateAt            int64                 `bson:"create_at"`
	Name                string                `bson:"name"`
	Type                int64                 `bson:"type"`
	LinkStructureGroups []*linkStructureGroup `bson:"link_structure_groups"`
}

type linkStructureRole struct {
	StructureId string   `bson:"structure_id"`
	Roles       []string `bson:"roles"`
}

type linkStructureGroup struct {
	StructureId string `bson:"structure_id"`
	BindGroupId string `bson:"bind_group_id"`
}

type role struct {
	Id           string `bson:"_id"`
	CreateUserId string `bson:"create_user_id"`
	CreateAt     int64  `bson:"create_at"`
	Name         string `bson:"name"`
	StructureId  string `bson:"structure_id"`
}

type functionGroup struct {
	Id           string `bson:"_id"`
	Name         string `bson:"name"`
	CreateUserId string `bson:"create_user_id"`
	CreateAt     int64  `bson:"create_at"`
	BindGroupId  string `bson:"bind_group_id"`
	Type         int64  `bson:"type"`
	StructureId  string `bson:"structure_id"`
}

type function struct {
	Id           string  `bson:"_id"`
	Name         string  `bson:"name"`
	Api          string  `bson:"api"`
	ApiTag       string  `bson:"api_tag"`
	Type         int64   `bson:"type"`
	CreateUserId string  `bson:"create_user_id"`
	CreateAt     int64   `bson:"create_at"`
	BindGroupId  string  `bson:"bind_group_id"`
	StructureId  string  `bson:"structure_id"`
	AuthTypes    []int64 `bson:"auth_types"`
	//authType container AuthTypeOfValcode. valTokenLife is access this function token expired time
	ValTokenLife   int64   `bson:"val_token_life"` //def: 0 second (your value must >= 60s)
	GrantPlatforms []int64 `bson:"grant_platforms"`
}

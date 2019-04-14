package permissionhandlers

const (
	dbName = "gs_permission"

	structureCollection = "structures"

	roleCollection = "roles"

	groupCollection = "groups"

	functionCollection = "functions"

	functionGroupCollection = "function_groups"
)

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
	Id           string `bson:"_id" json:"id"`
	SID          string `bson:"sid" json:"sid"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	AppId        string `bson:"app_id" json:"app_id"`
	Name         string `bson:"name" json:"name"`
	Type         int64  `bson:"type" json:"type"` //user or function structure
}

type structureState struct {
	AppId             string `json:"app_id"`
	FunctionStructure string `json:"function_structure"`
	UserStructure     string `json:"user_structure"`
	LastUpdateAt      int64  `json:"last_update_at"`
	LastUpdateBy      string `json:"last_update_by"`
}

type userOrientate struct {
	Id                  string                `bson:"_id" json:"id"`
	CreateAt            int64                 `bson:"create_at" json:"create_at"`
	UserId              string                `bson:"user_id" json:"user_id"`
	LinkStructureRoles  []*linkStructureRole  `bson:"link_structure_roles" json:"link_structure_roles"`
	LinkStructureGroups []*linkStructureGroup `bson:"link_structure_groups" json:"link_structure_groups"`
}

type userGroup struct {
	Id                  string                `bson:"_id" json:"id"`
	CreateUserId        string                `bson:"create_user_id" json:"create_user_id"`
	CreateAt            int64                 `bson:"create_at" json:"create_at"`
	Name                string                `bson:"name" json:"name"`
	Type                int64                 `bson:"type" json:"type"`
	LinkStructureGroups []*linkStructureGroup `bson:"link_structure_groups" json:"link_structure_groups"`
}

type linkStructureRole struct {
	StructureId string   `bson:"structure_id" json:"structure_id"`
	Roles       []string `bson:"roles" json:"roles"`
}

type linkStructureGroup struct {
	StructureId string `bson:"structure_id" json:"structure_id"`
	BindGroupId string `bson:"bind_group_id" json:"bind_group_id"`
}

type role struct {
	Id           string `bson:"_id" json:"id"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	Name         string `bson:"name" json:"name"`
	StructureId  string `bson:"structure_id" json:"structure_id"`
}

type functionGroup struct {
	Id           string `bson:"_id" json:"id"`
	Name         string `bson:"name" json:"name"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	BindGroupId  string `bson:"bind_group_id" json:"bind_group_id"`
	Type         int64  `bson:"type" json:"type"`
	StructureId  string `bson:"structure_id" json:"structure_id"`
}

type function struct {
	Id           string  `bson:"_id" json:"id"`
	SID          string  `bson:"sid" json:"sid"`
	Name         string  `bson:"name" json:"name"`
	Api          string  `bson:"api" json:"api"`
	Type         int64   `bson:"type" json:"type"`
	CreateUserId string  `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64   `bson:"create_at" json:"create_at"`
	BindGroupId  string  `bson:"bind_group_id" json:"bind_group_id"`
	StructureId  string  `bson:"structure_id" json:"structure_id"`
	AuthTypes    []int64 `bson:"auth_types" json:"auth_types"`
	//authType container AuthTypeOfValcode. valTokenLife is access this function token expired time
	ValTokenLife   int64   `bson:"val_token_life" json:"val_token_life"` //def: 0 second (your value must >= 60s)
	GrantPlatforms []int64 `bson:"grant_platforms" json:"grant_platforms"`
}

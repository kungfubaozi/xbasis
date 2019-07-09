package permissionhandlers

const (
	dbName = "xbs_permission"

	roleCollection = "user_roles"

	groupCollection = "user_groups"

	groupUsersCollection = "user_group_relation"

	functionCollection = "functions"

	functionGroupCollection = "function_groups"

	userRoleRelationCollection = "user_roles_relation"

	functionRoleRelationCollection = "function_roles_relation"

	functionIndex = "xbs-functions"

	functionGroupRelationIndex = "xbs-function-groups"

	roleIndex = "xbs-roles"

	statementIndex = "xbs-statement."
)

type accessibale struct {
	UserId     string `bson:"user_id" json:"user_id"`
	FunctionId string `bson:"function_id" json:"function_id"`
	AppId      string `bson:"app_id" json:"app_id"`
	RoleId     string `bson:"role_id" json:"role_id"`
	Recheck    bool   `json:"recheck"`
	Access     bool   `json:"access"`
}

//直接联系(用在es上)
type directrelation struct {
	RoleId     string `bson:"role_id" json:"roleId"`
	UserId     string `bson:"user_id" json:"userId"`
	FunctionId string `bson:"function_id" json:"functionId"`
	User       bool   `bson:"user" json:"user"`
	Function   bool   `bson:"function" json:"function"`
	Enabled    bool   `bson:"enabled" json:"enabled"`
}

type DurationAccessCredential struct {
	FromClientId string
	RefClientId  string
	FuncId       string
	Timestamp    int64
	FromAuth     bool
	AppId        string
}

type DurationAccessToken struct {
	ClientId string
	FuncId   string
	User     string
	Times    int64
	MaxTimes int64
	Auth     bool
	From     string
}

type durationAccess struct {
	User          string
	From          string
	FuncId        string
	Life          int64
	ClientId      string
	CreateAt      int64
	Stat          string
	Code          int64
	CodeExpiredAt int64
	Key           string
	Auth          bool
}

type userRolesRelation struct {
	CreateAt int64    `bson:"create_at" json:"create_at"`
	UserId   string   `bson:"user_id" json:"user_id"`
	Roles    []string `bson:"roles" json:"roles"`
	AppId    string   `bson:"app_id" json:"app_id"`
}

type functionRolesRelation struct {
	CreateAt   int64    `bson:"create_at" json:"create_at"`
	FunctionId string   `bson:"function_id" json:"function_id"`
	Roles      []string `bson:"roles" json:"roles"`
	AppId      string   `bson:"app_id" json:"app_id"`
}

type userGroupsRelation struct {
	CreateAt    int64    `bson:"create_at" json:"create_at"`
	UserId      string   `bson:"user_id" json:"user_id"`
	AppId       string   `bson:"app_id" json:"app_id"`
	BindGroupId []string `bson:"bind_group_id" json:"bind_group_id"` //用户在同一结构下可能会在多个组内
}

type userGroup struct {
	Id           string `bson:"_id" json:"id"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	Name         string `bson:"name" json:"name"`
	Type         int64  `bson:"type" json:"type"`
	AppId        string `bson:"app_id" json:"app_id"`
	BindGroupId  string `bson:"bind_group_id" json:"bind_group_id"`
}

type role struct {
	Id           string `bson:"_id" json:"id"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	Name         string `bson:"name" json:"name"`
	AppId        string `bson:"app_id" json:"app_id"`
}

type functionGroup struct {
	Id           string `bson:"_id" json:"id"`
	SID          string `bson:"sid" json:"sid"`
	Name         string `bson:"name" json:"name"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	BindGroupId  string `bson:"bind_group_id" json:"bind_group_id"`
	Type         int64  `bson:"type" json:"type"`
	AppId        string `bson:"app_id" json:"app_id"`
}

type function struct {
	Id           string  `bson:"_id" json:"id"`
	SID          string  `bson:"sid" json:"sid"`
	Desc         string  `bson:"desc" json:"desc"`
	Name         string  `bson:"name" json:"name" es:"not_analyzed"`
	Api          string  `bson:"api" json:"api"`
	Type         int64   `bson:"type" json:"type"`
	CreateUserId string  `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64   `bson:"create_at" json:"create_at"`
	BindGroupId  string  `bson:"bind_group_id" json:"bind_group_id"`
	AppId        string  `bson:"app_id" json:"app_id"`
	AuthTypes    []int64 `bson:"auth_types" json:"auth_types"`
	//authType container AuthTypeOfValcode. valTokenLife is access this function token expired time
	ValTokenTimes  int64   `bson:"val_token_times" json:"val_token_times"` //可以使用的次数 >=1
	GrantPlatforms []int64 `bson:"grant_platforms" json:"grant_platforms"`
	//representation validation does not require judging the application to which it belongs, and each application can share this function (roles need to be set to null)
	Share bool `bson:"share" json:"share"`
}

type SimplifiedFunction struct {
	Id             string  `json:"id"`
	AuthTypes      []int64 `json:"auth_types"`
	ValTokenTimes  int64   `bson:"val_token_times" json:"val_token_times"` //可以使用的次数 >=1
	GrantPlatforms []int64 `json:"grant_platforms"`
	Share          bool    `bson:"share" json:"share"`
	AppId          string  `bson:"app_id" json:"app_id"`
	Path           string  `bson:"path" json:"path"`
	Name           string  `bson:"name" json:"name"`
}

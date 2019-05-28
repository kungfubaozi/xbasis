package applicationhanderls

const (
	dbName = "gs_applications"

	applicationCollection = "applications"
)

type appInfo struct {
	Id           string        `bson:"_id" json:"id"`
	SID          string        `bson:"sid" json:"sid"`
	Name         string        `bson:"name" json:"name"`
	Desc         string        `bson:"desc" json:"desc"`
	CreateUserId string        `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64         `bson:"create_at" json:"create_at"`
	Settings     *appSetting   `bson:"settings" json:"settings"`
	Clients      []*appClient  `bson:"clients" json:"clients"`
	Type         int64         `bson:"type" json:"type"` //main application, provider sso, permission, user and so more
	UserS        *appStructure `bson:"user_s" json:"user_s"`
	FunctionS    *appStructure `bson:"function_s" json:"function_s"`
}

type appSetting struct {
	Enabled       int64                 `bson:"enabled" json:"enabled"`
	SyncUserURL   string                `bson:"sync_user_url" json:"sync_user_url"` //sync new user to your application database
	RedirectURL   string                `bson:"redirect_url" json:"redirect_url"`
	Quarantine    int64                 `bson:"quarantine" json:"quarantine"`
	MustSync      bool                  `bson:"must_sync" json:"must_sync"`
	AllowNewUsers *allowNewUsersToEnter `bson:"allow_new_users" json:"allow_new_users"` //允许新用户登录，如果允许那么会把默认的structure下的group/role设置给新进入的用户
}

type allowNewUsersToEnter struct {
	Enabled          bool     `bson:"enabled" json:"enabled"`
	DefaultStructure string   `bson:"default_structure" json:"default_structure"`
	DefaultGroup     string   `bson:"default_group" json:"default_group"`
	DefaultRole      []string `bson:"default_role" json:"default_role"`
}

type appClient struct {
	Id       string `bson:"id" json:"id"`
	Platform int64  `bson:"platform" json:"platform"`
	Enabled  int64  `bson:"enabled" json:"enabled"`
}

type appStructure struct {
	Id           string `bson:"id" json:"id"`
	LastUpdateAt int64  `bson:"last_update_at" json:"last_update_at"`
	LastUpdateBy string `bson:"last_update_by" json:"last_update_by"`
}

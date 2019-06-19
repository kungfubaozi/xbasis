package applicationhanderls

const (
	dbName = "gs_applications"

	applicationCollection = "applications"
)

type appInfo struct {
	Id           string       `bson:"_id" json:"id"`
	SID          string       `bson:"sid" json:"sid"`
	Name         string       `bson:"name" json:"name"`
	Desc         string       `bson:"desc" json:"desc"`
	CreateUserId string       `bson:"create_user_id" json:"create_user_id"`
	CreateAt     int64        `bson:"create_at" json:"create_at"`
	Settings     *appSetting  `bson:"settings" json:"settings"`
	Clients      []*appClient `bson:"clients" json:"clients"`
	Type         int64        `bson:"type" json:"type"`             //main application, provider sso, permission, user and so more
	SecretKey    string       `json:"secret_key" json:"secret_key"` //密钥，主要是用来加密传输数据
}

type appSetting struct {
	Enabled     int64  `bson:"enabled" json:"enabled"`
	SyncUserURL string `bson:"sync_user_url" json:"sync_user_url"` //sync new user to your application database
	RedirectURL string `bson:"redirect_url" json:"redirect_url"`
	Quarantine  bool   `bson:"quarantine" json:"quarantine"` //不开放注册
	//如果需要授权，会把用户信息同步至SyncUserURL（不为nil的情况下）
	Authorize     bool                  `bson:"authorize" json:"authorize"`             //进入时是否需要授权
	AllowNewUsers *allowNewUsersToEnter `bson:"allow_new_users" json:"allow_new_users"` //允许新用户登录，如果允许那么会把默认的structure下的group/role设置给新进入的用户
}

type allowNewUsersToEnter struct {
	Enabled      bool     `bson:"enabled" json:"enabled"`
	DefaultGroup string   `bson:"default_group" json:"default_group"`
	DefaultRole  []string `bson:"default_role" json:"default_role"`
}

type appClient struct {
	Id       string `bson:"id" json:"id"`
	Platform int64  `bson:"platform" json:"platform"`
	Enabled  int64  `bson:"enabled" json:"enabled"`
}

type syncLog struct {
	UserId      string `json:"user_id"`
	AppId       string `json:"app_id"`
	SHARelation string `json:"sha_relation"`
	Timestamp   int64  `json:"timestamp"`
}

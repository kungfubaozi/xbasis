package applicationhanderls

const (
	dbName = "gs_application"

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
	Main         int64         `bson:"main" json:"main"` //main application, provider sso, permission, user and so more
	UserS        *appStructure `bson:"user_s" json:"user_s"`
	FunctionS    *appStructure `bson:"function_s" json:"function_s"`
}

type appSetting struct {
	Enabled     int64  `bson:"enabled" json:"enabled"`
	SyncUserURL string `bson:"sync_user_url" json:"sync_user_url"` //sync new user to your application database
	OpenMode    int64  `bson:"open_mode" json:"open_mode"`         //open mode
	RedirectURL string `bson:"redirect_url" json:"redirect_url"`
	Quarantine  int64  `bson:"quarantine" json:"quarantine"` //create local self database
	MustSync    bool   `bson:"must_sync" json:"must_sync"`
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

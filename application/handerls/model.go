package applicationhanderls

const (
	dbName = "gs_application"

	applicationCollection = "applications"
)

type appInfo struct {
	Id           string       `bson:"_id"`
	Name         string       `bson:"name"`
	Desc         string       `bson:"desc"`
	CreateUserId string       `bson:"create_user_id"`
	CreateAt     int64        `bson:"create_at"`
	Settings     *appSetting  `bson:"settings"`
	Clients      []*appClient `bson:"clients"`
	Main         int64        `bson:"main"` //main application, provider sso, permission, user and so more
}

type appSetting struct {
	Enabled     int64  `bson:"enabled"`
	SyncUserURL string `bson:"sync_user_url"` //sync new user to your application database
	OpenMode    int64  `bson:"open_mode"`     //open mode
	RedirectURL string `bson:"redirect_url"`
	Quarantine  int64  `bson:"quarantine"` //create local self database
}

type appClient struct {
	Id       string `bson:"id"`
	Platform int64  `bson:"platform"`
	Enabled  int64  `bson:"enabled"`
}

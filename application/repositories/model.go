package application_repositories

type AppInfo struct {
	Id           string       `bson:"_id"`
	Name         string       `bson:"name"`
	Desc         string       `bson:"desc"`
	CreateUserId string       `bson:"create_user_id"`
	CreateAt     int64        `bson:"create_at"`
	Settings     *AppSetting  `bson:"settings"`
	Clients      []*AppClient `bson:"clients"`
}

type AppSetting struct {
	Enabled     int64  `bson:"enabled"`
	SyncUserURL string `bson:"sync_user_url"` //sync new user to your application database
	OpenMode    int64  `bson:"open_mode"`     //open mode
}

type AppClient struct {
	Id       string `bson:"id"`
	Secret   string `bson:"secret"`
	Platform int64  `bson:"platform"`
	Enabled  int64  `bson:"enabled"`
}

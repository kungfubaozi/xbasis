package safetyhanders

const (
	dbName = "gs_safety"

	blacklistCollection = "blacklists"
)

type blacklist struct {
	Type         int64  `bson:"type"`
	Content      string `bson:"content"`
	CreateAt     int64  `bson:"create_at"`
	CreateUserId string `bson:"create_user_id"`
}

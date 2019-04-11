package safetyhanders

const (
	dbName = "gs_safety"

	blacklistCollection = "blacklists"
)

type blacklist struct {
	Type         int64  `bson:"type" json:"type"`
	Content      string `bson:"content" json:"content"`
	CreateAt     int64  `bson:"create_at" json:"create_at"`
	CreateUserId string `bson:"create_user_id" json:"create_user_id"`
}

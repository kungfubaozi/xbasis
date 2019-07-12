package xbasistransport

type AppFunction struct {
	Id               string
	AuthTypes        []int64
	Name             string
	ValTokenTimes    int64
	NoGrantPlatforms []int64
	Share            bool
	AppId            string
	Path             string
	Version          int64
	Unavailable      bool
}

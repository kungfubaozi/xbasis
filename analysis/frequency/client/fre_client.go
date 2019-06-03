package frequencyclient

type Frequency struct {
	Action    string `json:"action"`
	Tag       string `json:"tag"`
	Timestamp int64  `json:"timestamp"`
}

func NewClient() {

}

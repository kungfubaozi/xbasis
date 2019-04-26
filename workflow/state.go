package workflow

type state struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

var (
	Success = &state{Ok: true, Message: "ok", Code: 0}
)

package types

type HttpRequestMethod int64

const (
	GetRequest HttpRequestMethod = iota

	PostRequest

	PutRequest

	DeleteRequest
)

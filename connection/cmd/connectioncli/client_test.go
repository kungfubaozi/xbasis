package connectioncli

import (
	"konekko.me/gosion/message"
	"testing"
)

func TestNewClient(t *testing.T) {

	content := make(map[string]interface{})
	content["key"] = &message.Message{
		To:      "asdfasd",
		Content: "asdfasdf",
	}

	op, err := NewClient()
	if err != nil {
		panic(err)
	}
	err = op.BroadcastAll(&Message{
		Content: content,
	})
	if err != nil {
		panic(err)
	}

}

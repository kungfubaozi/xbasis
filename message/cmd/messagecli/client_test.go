package messagecli

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		panic(err)
	}
	c.SendSMS("13222021207", "this test message")
}

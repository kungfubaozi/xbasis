package permissionhandlers

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"testing"
	"time"
)

func TestDat(t *testing.T) {
	credential := &durationAccessCredential{
		FromClientId: "81892345",
		RefClientId:  "3523454395",
		FuncId:       "12938490138459023485092359038",
		Timestamp:    time.Now().UnixNano(),
	}

	b, err := msgpack.Marshal(credential)
	if err != nil {
		panic(errstate.ErrSystem)
	}

	c, err := encrypt.AESEncrypt(b, []byte(encrypt.Md5("234324")))
	if err != nil {
		panic(errstate.ErrSystem)
	}

	fmt.Println("data", c)
}

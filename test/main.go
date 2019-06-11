package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/errstate"
	"time"
)

type durationAccessCredential struct {
	FromClientId string
	RefClientId  string
	FuncId       string
	Timestamp    int64
	FromAuth     bool
}

func main() {
	//fmt.Println(gs_commons_constants.AuthTypeOfToken, gs_commons_constants.AuthTypeOfFace,
	//	gs_commons_constants.AuthTypeOfMobileConfirm, gs_commons_constants.AuthTypeOfValcode,
	//	gs_commons_constants.AuthTypeOfMiniProgramCodeConfirm)
	//a := []int64{2, 3, 4, 6, 9}
	//size := len(a)
	//s := size / 2
	//fmt.Println(s, size-s)
	//fmt.Println(a[:s], a[s:])
	//fmt.Println(gs_commons_generator.NewIDG().Get())
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

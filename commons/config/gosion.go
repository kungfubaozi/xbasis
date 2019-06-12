package gs_commons_config

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/constants"
)

type GosionConfiguration struct {
	AccessTokenExpiredTime           int64  //def: 10min unit:min
	RefreshTokenExpiredTime          int64  //def: 7day unit:day
	EmailValidateTemplate            string //no def
	EmailVerificationCodeExpiredTime int64  //def: 10*60s unit:second
	PhoneValidateTemplate            string //no def
	PhoneVerificationCodeExpiredTime int64  //no 10*60s unit:second
	SendVerificationCodeType         int64
	//How long does the user log in interval start locking
	LoginIntervalToStartLock          int64 //def: 30days unit:day
	CurrencySecretKey                 string
	TokenSecretKey                    string
	RegisterType                      int64 //def: 1001 , [1001(phone), 1002(email), 1003(face)]
	LoginType                         int64 //def: all
	DurationAccessTokenRetryTime      int64 //def: 60s
	DurationAccessTokenSendCodeToType int64 //def: 1001
}

type OnGosionConfigurationChanged func(config *GosionConfiguration)

func WatchGosionConfig(event OnGosionConfigurationChanged) {
	c := NewConnect("192.168.2.57:2181")
	//one same service process
	acl := zk.WorldACL(zk.PermAll)
	_, err := c.Create(gs_commons_constants.GosionConfiguration, nil, 1, acl)
	if err != nil {
		fmt.Println("node register error:", err)
		//return
	}

	back := func(b []byte) bool {
		var config GosionConfiguration
		err := msgpack.Unmarshal(b, &config)
		if err != nil {
			return true
		}
		fmt.Println("receiver new config")
		event(&config)
		return true
	}

	v, _, err := c.Get(gs_commons_constants.GosionConfiguration)
	if err != nil {
		fmt.Println("err", err)
	} else {
		if v != nil {
			back(v)
		} else {
			fmt.Println("nothing")
		}
	}

	watch(c, gs_commons_constants.GosionConfiguration, func(data []byte, version int32) bool {
		return back(data)
	})
}

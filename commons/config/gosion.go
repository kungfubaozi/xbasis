package gs_commons_config

import (
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
		return
	}

	watch(c, gs_commons_constants.GosionConfiguration, func(data []byte, version int32) bool {
		var config GosionConfiguration
		err := msgpack.Unmarshal(data, &config)
		if err != nil {
			return false
		}
		event(&config)
		return false
	})
}

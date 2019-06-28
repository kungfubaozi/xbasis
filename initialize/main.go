package main

import (
	"flag"
	"fmt"
	"github.com/vmihailenco/msgpack"
	xconfig "konekko.me/xbasis/commons/config"
	constants "konekko.me/xbasis/commons/constants"
	"konekko.me/xbasis/commons/encrypt"
	generator "konekko.me/xbasis/commons/generator"
	regx "konekko.me/xbasis/commons/regx"
	"os"
	"strconv"
	"time"
)

var (
	//facePath   string
	email      string
	phone      string
	account    string
	username   string
	name       string
	enterprise string
	config     string
	desc       string
)

func init() {
	//flag.StringVar(&facePath, "face", "system_admin.jpg", "def admin face local path.")
	flag.StringVar(&email, "email", "", "def admin email.")
	flag.StringVar(&phone, "phone", "", "def admin phone.")
	flag.StringVar(&username, "username", "Admin", "def admin username.")
	flag.StringVar(&name, "name", "admin", "def admin realName.")
	flag.StringVar(&enterprise, "enterprise", "", "your enterprise name.")
	flag.StringVar(&desc, "desc", "", "your enterprise desc.")
	flag.StringVar(&config, "config", "", "zookeeper config address.")
	flag.StringVar(&account, "account", "admin", "def admin account.")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	//if len(facePath) == 0 {
	//	fmt.Println("please set admin face local path.")
	//	return
	//}
	if !regx.Email(email) {
		fmt.Println("email format err.")
		return
	}
	if !regx.Phone(phone) {
		fmt.Println("phone format err.")
		return
	}
	if len(username) < 2 {
		fmt.Println("admin username length must >= 2")
		return
	}
	if len(account) < 2 {
		fmt.Println("admin account length must >= 2")
		return
	}
	if len(name) < 2 {
		fmt.Println("admin name length must >= 2")
		return
	}
	if len(enterprise) < 2 {
		fmt.Println("please set your enterprise name.")
		return
	}
	if len(config) < 7 {
		fmt.Println("please set zookeeper config server address.")
		return
	}
	fmt.Println("system initialize...")
	time.Sleep(200)

	id := generator.NewIDG()

	secretKey := encrypt.Md5(time.Now().String())

	initConfig := &xconfig.GosionInitializeConfig{
		AppName:       enterprise,
		UserId:        id.Get(),
		Desc:          desc,
		Username:      username,
		Phone:         phone,
		Email:         email,
		SecretKey:     secretKey,
		UserAppId:     id.Short(),
		UserAppRoleId: id.UUID(),
		RouteAppId:    id.Short(),
		SafeAppId:     id.Short(),
		AdminAppId:    id.Short(),
	}

	configuration := &xconfig.GosionConfiguration{
		AccessTokenExpiredTime:            10 * 60 * 1000,
		RefreshTokenExpiredTime:           7 * 24 * 60 * 60 * 1000,
		EmailVerificationCodeExpiredTime:  10 * 60 * 1000,
		PhoneVerificationCodeExpiredTime:  10 * 60 * 1000,
		LoginIntervalToStartLock:          30 * 24 * 60 * 60 * 1000,
		CurrencySecretKey:                 encrypt.Md5("currency-secret" + strconv.FormatInt(time.Now().UnixNano(), 10)),
		RegisterType:                      1001,
		LoginType:                         1001,
		TokenSecretKey:                    secretKey,
		DurationAccessTokenSendCodeToType: 1001,
		DurationAccessTokenRetryTime:      60,
	}

	b, err := msgpack.Marshal(initConfig)
	if err != nil {
		panic(err)
	}

	if len(b) > 0 {
		fmt.Println(string(b))

	}

	//once
	c := xconfig.NewConnect(config)
	//set def configs
	_, s, err := c.Get(constants.ZKWatchInitializeConfigPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("version", s.Version)
	v := 0
	if s.DataLength == 0 && s.Version == 0 {
		v = 0
		fmt.Println("r")
	} else {
		v = int(s.Version)
	}

	fmt.Println("version -", s.Version)

	_, err = c.Set(constants.ZKWatchInitializeConfigPath, b, int32(v))
	if err != nil {
		fmt.Println("set init config err", err)
	}

	_, s, err = c.Get(constants.GosionConfiguration)
	if err != nil {

	}
	fmt.Println("version", s.Version)
	v = 0
	if s.DataLength == 0 && s.Version == 0 {
		v = 0
		fmt.Println("r")
	} else {
		v = int(s.Version)
	}

	fmt.Println("version -", s.Version)
	b, err = msgpack.Marshal(configuration)
	if err != nil {
		panic(err)
	}
	_, err = c.Set(constants.GosionConfiguration, b, int32(v))
	if err != nil {
		panic(err)
	}

}

func usage() {
	fmt.Fprintf(os.Stderr, `Gosion initlizate command
Options:
`)
	flag.PrintDefaults()
}

package main

import (
	"flag"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/encrypt"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/regx"
	"os"
	"strconv"
	"time"
)

var (
	//facePath   string
	email      string
	phone      string
	username   string
	name       string
	password   string
	enterprise string
	config     string
	desc       string
)

func init() {
	//flag.StringVar(&facePath, "face", "system_admin.jpg", "def admin face local path.")
	flag.StringVar(&email, "email", "", "def admin email.")
	flag.StringVar(&phone, "phone", "", "def admin phone.")
	flag.StringVar(&username, "username", "admin", "def admin username.")
	flag.StringVar(&name, "name", "admin", "def admin realName.")
	flag.StringVar(&password, "password", "admin123", "def admin password.")
	flag.StringVar(&enterprise, "enterprise", "", "your enterprise name.")
	flag.StringVar(&desc, "desc", "", "your enterprise desc.")
	flag.StringVar(&config, "config", "", "zookeeper config address.")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	//if len(facePath) == 0 {
	//	fmt.Println("please set admin face local path.")
	//	return
	//}
	if !gs_commons_regx.Email(email) {
		fmt.Println("email format err.")
		return
	}
	if !gs_commons_regx.Phone(phone) {
		fmt.Println("phone format err.")
		return
	}
	if len(username) < 2 {
		fmt.Println("admin username length must >= 2")
		return
	}
	if len(name) < 2 {
		fmt.Println("admin name length must >= 2")
		return
	}
	if len(password) < 6 {
		fmt.Println("admin password length must >= 6")
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

	id := gs_commons_generator.NewIDG()

	initConfig := &gs_commons_config.GosionInitializeConfig{
		AppId:       id.Short(),
		AppName:     enterprise,
		UserId:      id.Get(),
		Desc:        desc,
		Username:    username,
		Email:       email,
		Password:    password,
		WebClientId: id.Short(),
	}

	fmt.Println("init application def clientId(web) is:", initConfig.WebClientId)

	configuration := &gs_commons_config.GosionConfiguration{
		AccessTokenExpiredTime:           10 * 60 * 1000,
		RefreshTokenExpiredTime:          7 * 24 * 60 * 60 * 1000,
		EmailVerificationCodeExpiredTime: 10 * 60 * 1000,
		PhoneVerificationCodeExpiredTime: 10 * 60 * 1000,
		LoginIntervalToStartLock:         30 * 24 * 60 * 60 * 1000,
		CurrencySecretKey:                gs_commons_encrypt.Md5("currency-secret" + initConfig.AppId + strconv.FormatInt(time.Now().UnixNano(), 10)),
		RegisterType:                     1001 | 1002 | 1003,
		LoginType:                        1001 | 1002 | 1003,
		TokenSecretKey:                   gs_commons_encrypt.Md5(time.Now().String()),
	}

	b, err := msgpack.Marshal(initConfig)
	if err != nil {
		panic(err)
	}

	//once
	c := gs_commons_config.NewConnect(config)
	//set def configs
	_, err = c.Set(gs_commons_constants.ZKWatchInitializeConfigPath, b, 0)
	if err != nil {
		fmt.Println("set init config err", err)
	}

	b, err = msgpack.Marshal(configuration)
	if err != nil {
		panic(err)
	}
	_, err = c.Set(gs_commons_constants.GosionConfiguration, b, 4)
	if err != nil {
		panic(err)
	}
	//
	//v, _, err := c.Get(gs_commons_constants.GosionConfiguration)
	//if err != nil {
	//	fmt.Println("err", err)
	//} else {
	//	if v != nil {
	//		var config gs_commons_config.GosionConfiguration
	//		err := msgpack.Unmarshal(b, &config)
	//		if err != nil {
	//			panic(err)
	//		}
	//		fmt.Println("init config ok", config)
	//	} else {
	//		fmt.Println("nothing")
	//	}
	//}

}

func usage() {
	fmt.Fprintf(os.Stderr, `Gosion initlizate command
Options:
`)
	flag.PrintDefaults()
}

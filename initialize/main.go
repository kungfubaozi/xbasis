package main

import (
	"flag"
	"fmt"
	"github.com/vmihailenco/msgpack"
	"golang.org/x/crypto/bcrypt"
	"konekko.me/gosion/commons/config"
	"konekko.me/gosion/commons/constants"
	"konekko.me/gosion/commons/generator"
	"konekko.me/gosion/commons/regx"
	"os"
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

	node := gs_commons_generator.ID()

	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	config := &gs_commons_config.GosionInitializeConfig{
		AppId:    node.Generate().String(),
		AppName:  enterprise,
		UserId:   node.Generate().String(),
		Desc:     desc,
		Username: username,
		Password: string(b),
	}

	b, err = msgpack.Marshal(config)
	if err != nil {
		panic(err)
	}

	//once
	c := gs_commons_config.NewConnect("192.168.2.57:2181")
	//set def configs
	_, err = c.Set(gs_commons_constants.ZKWatchInitializeConfigPath, b, 0)
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

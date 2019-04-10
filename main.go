package main

import (
	"errors"
	"fmt"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"reflect"
)

type Message struct {
	UserId   string
	TargetId string
	Type     int64
	Content  map[string]interface{}
}

func main() {
	//consulConfig := api.DefaultConfig()
	//consulConfig.Address = "192.168.80.67:8500"
	//
	//client, err := api.NewClient(consulConfig)
	//if err != nil {
	//	panic(err)
	//}
	//
	//pair, _, err := client.KV().Get("testvalue", nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("value is ", string(pair.Value))
	//
	//c, _, err := zk.Connect([]string{"192.168.2.57:2181"}, time.Second) //*10)
	//if err != nil {
	//	panic(err)
	//}
	//acl := zk.WorldACL(zk.PermAll)
	//_, err = c.Create("/gosion.test.3", nil, int32(0), acl)
	//if err != nil {
	//	fmt.Printf("create: %+v\n", err)
	//}
	//
	//for {
	//	_, _, ch, err := c.GetW("/gosion.test.3")
	//	if err != nil {
	//		panic(err)
	//	}
	//	select {
	//	case e := <-ch:
	//		if e.Type == zk.EventNodeDataChanged {
	//			v, s, err := c.Get("/gosion.test.3")
	//			if err != nil {
	//				fmt.Println("err", err)
	//			} else {
	//				var config gs_commons_config.GosionInitializeConfig
	//				err = msgpack.Unmarshal(v, &config)
	//				if err == nil {
	//					fmt.Println(config.UserId)
	//					return
	//				}
	//				fmt.Println("version", s.Version)
	//			}
	//
	//		}
	//	}
	//}

	//s := &UserSecretKey{}
	////s.SecretKey = make(chan string)
	//t := test(s)
	//s.SecretKey = "324324"
	//t()
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	s.SecretKey = "2343324"
	//	wg.Done()
	//}()
	//wg.Wait()
	//
	////fmt.Println(s.SecretKey)
	//
	//t()
	//dat := &permission_repositories.DurationAccess{
	//	ClientId: "192834-09123840989080918240-asdf",
	//}

	//str := "75089174091823740189237498"
	////
	////b, err := msgpack.Marshal(dat)
	////if err != nil {
	////	panic(err)
	////}
	//
	//for i := 0; i <= 100; i++ {
	//	v, err := gs_commons_encrypt.AESEncrypt([]byte(str), []byte("912ec803b2ce49e4a541068d495ab570"))
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(v)
	//}

	//fmt.Println(time.Now().UnixNano())

	//var createAt int64
	//createAt = 1553175979844167000
	//
	//fmt.Println(time.Now().UnixNano()/1e9-createAt/1e9 >= 60)

	//fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000))

	//s := time.Now().UnixNano()
	//out := new(gs_commons_dto.Status)
	//Add(out)
	//v := time.Now().UnixNano() - s
	//fmt.Println(out.State)
	//fmt.Println(v / 1e6)

	//id := gs_commons_generator.NewIDG()

	//fmt.Println(id.Short())
	//
	//k := gs_commons_encrypt.Md5(id.Get() + id.Short())
	//
	//fmt.Println(k)
	//
	//v, err := gs_commons_encrypt.AESEncrypt([]byte("this is test data"), []byte(k))
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(v))

	//
	//sut := &authentication_repositories.SimpleUserToken{
	//	UserId:   id.Get(),
	//	AppId:    id.Get(),
	//	ClientId: id.Get(),
	//	Relation: id.Get(),
	//}
	//
	//fmt.Println(sut)
	//for i := 0; i < 100; i++ {
	//	fmt.Println(id.Short(), gs_commons_map.Hashcode(id.Short()))
	//}
	//
	//fmt.Println(gs_commons_map.Hashcode("tvrkak9uttbova=="))
	//
	//b, err := msgpack.Marshal(sut)
	//if err != nil {
	//	panic(err)
	//}
	//
	//s, err := gs_commons_encrypt.AESEncrypt(b, []byte("912ec803b2ce49e4a541068d495ab570"))
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("value", id.Node().Generate().Base58())
	//
	//key := id.Short() + "." + id.Get()
	//fmt.Println("value", key[strings.Index(key, ".")+1:])

	//es := make(chan *gs_commons_dto.State)
	//
	//for {
	//	select {
	//	case e := <-es:
	//		fmt.Println("e", e)
	//		break
	//	default:
	//		//fmt.Println("out")
	//		break
	//	}
	//}

	fmt.Println(4 << 6)

}

type TestController func()

func test(config *UserSecretKey) TestController {
	return func() {
		fmt.Println(config.SecretKey)
	}
}

type UserSecretKey struct {
	SecretKey string
}

func Add(out *gs_commons_dto.Status) error {
	return ContextToAuthorize(out, func() *gs_commons_dto.State {
		test1(out)
		return nil
	})
}

func test1(out *gs_commons_dto.Status) {
	//out.State = errstate.ErrSystem
}

type ContextWrapperListener func() *gs_commons_dto.State

func ContextToAuthorize(out interface{}, listener ContextWrapperListener) error {
	s := reflect.ValueOf(out).Elem().FieldByName("State")
	if !s.CanSet() {
		return errors.New("no found 'State' filed")
	}
	state := s.Interface().(*gs_commons_dto.State)
	if state == nil {
		//state = new(gs_commons_dto.State)
	}
	s.Set(reflect.ValueOf(state))
	listener()
	if s.IsNil() {
		s.Set(reflect.ValueOf(errstate.ErrRequest))
		fmt.Println("err")
	}
	return nil
}

package main

import (
	"context"
	"errors"
	"fmt"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/permission/pb"
	"reflect"
	"sync"
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

	s := &UserSecretKey{}
	//s.SecretKey = make(chan string)
	t := test(s)
	s.SecretKey = "324324"
	t()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		s.SecretKey = "2343324"
		wg.Done()
	}()
	wg.Wait()

	//fmt.Println(s.SecretKey)

	t()
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

func Add(ctx context.Context, in *gs_service_permission.RoleRequest, out *gs_commons_dto.Status) error {
	return ContextToAuthorize(ctx, out, func() *gs_commons_dto.State {
		return nil
	})
}

type ContextWrapperListener func() *gs_commons_dto.State

func ContextToAuthorize(ctx context.Context, out interface{}, listener ContextWrapperListener) error {
	s := reflect.ValueOf(out).Elem().FieldByName("State")
	if !s.CanSet() {
		return errors.New("no found 'State' filed")
	}
	state := s.Interface().(*gs_commons_dto.State)
	if state == nil {
		state = new(gs_commons_dto.State)
	}
	s.Set(reflect.ValueOf(state))
	state = listener()
	return nil
}

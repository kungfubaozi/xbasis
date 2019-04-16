package authenticationhandlers

import (
	"fmt"
	"github.com/twinj/uuid"
	"konekko.me/gosion/commons/encrypt"
	"strconv"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {

	k := encrypt.Md5("currency-secret" + strconv.FormatInt(time.Now().UnixNano(), 10))

	m := &simpleUserToken{
		UserId:    uuid.NewV4().String(),
		AppId:     uuid.NewV4().String(),
		ClientId:  uuid.NewV4().String(),
		Relation:  uuid.NewV4().String(),
		Structure: uuid.NewV4().String(),
	}

	token, err := encodeToken(k,
		time.Hour*24*7, m)

	if err != nil {
		panic(err)
	}

	fmt.Println(token)

	fmt.Println(k)

}

func TestDecode(t *testing.T) {
	claims, err := decodeToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJUb2tlbiI6eyJVc2VySWQiOiI0ZjMzYzg2Yy1iYjhiLTQyYmQtOTQ2YS05ZmFmYTk0Njg4NmQiLCJBcHBJZCI6IjEwYzNlZjAwLWExYzEtNDA4MC04ZDlhLTVkZGM2NTAzOGQ1OSIsIkNsaWVudElkIjoiNjVkZDc3MjctY2RmMS00ZjNlLWEwMzMtMzk4Zjc5N2UwODM4IiwiUmVsYXRpb24iOiI1NzVkODU4NC1mMTA1LTQ2YTQtOGNhMy1lYzJmZjAzOTNmZGQiLCJUeXBlIjowLCJTdHJ1Y3R1cmUiOiJjYjA2MWVmMC02Y2YwLTQ4MjItYWNjNy0xODhlOTYyMzVjZmMifSwiZXhwIjoxNTU2MDExMDEwLCJpYXQiOjE1NTU0MDYyMTA3Mjc1NzYwMDAsImlzcyI6Imdvc2lvbi5hdXRoZW50aWNhdGUifQ.wyrH7ivjLK-TJZXPZ4TAXK-ADH7niOyZiS0TwGfdbzo", "0d81fedacc2cc874e8ae336b808b70ee")
	if err != nil {
		panic(err)
	}

	fmt.Println(claims.IssuedAt)

	fmt.Println(time.Now().UnixNano()-claims.IssuedAt <= 2*60*1e9)
	fmt.Println(claims.ExpiresAt, time.Now().Unix())
}

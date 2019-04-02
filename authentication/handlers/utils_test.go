package authenticationhandlers

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/twinj/uuid"
	"github.com/vmihailenco/msgpack"
	"konekko.me/gosion/authentication/repositories"
	"konekko.me/gosion/commons/encrypt"
	"strconv"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {

	//refreshToken, err := encodeToken(gs_commons_encrypt.Md5("currency-secret"+strconv.FormatInt(time.Now().UnixNano(), 10)),
	//	time.Hour*24*7, )
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(refreshToken)

	k := gs_commons_encrypt.Md5("currency-secret" + strconv.FormatInt(time.Now().UnixNano(), 10))

	//h := sha1.New()
	//h.Write([]byte(b))
	//v := fmt.Sprintf("%x", h.Sum(nil))
	s := time.Now().Second()
	for i := 0; i < 3; i++ {
		m := &authentication_repositories.SimpleUserToken{
			UserId:    uuid.NewV4().String(),
			AppId:     uuid.NewV4().String(),
			ClientId:  uuid.NewV4().String(),
			Relation:  uuid.NewV4().String(),
			Structure: uuid.NewV4().String(),
		}
		b, err := msgpack.Marshal(m)
		if err != nil {
			panic(err)
		}

		h := hmac.New(sha256.New, []byte(k))
		h.Write(b)

		v := uuid.New(h.Sum(nil)).String()

		fmt.Println(v)

	}
	fmt.Println(fmt.Sprintf("%d %d", s, time.Now().Second()))
}

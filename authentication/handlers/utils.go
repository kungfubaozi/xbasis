package authenticationhandlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"konekko.me/xbasis/commons/dto"
	"konekko.me/xbasis/commons/errstate"
	"konekko.me/xbasis/connection/cmd/connectioncli"
	"strings"
	"time"
)

type UserClaims struct {
	Token *simpleUserToken

	jwt.StandardClaims
}

func decodeToken(token, tokenKey string) (*UserClaims, error) {
	if len(token) <= 128 {
		return nil, errors.New("err token")
	}
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*UserClaims); ok {
		return claims, nil
	}
	return nil, err
}

//加密token
func encodeToken(tokenKey string, et time.Duration, authorize *simpleUserToken) (string, error) {
	expireTime := time.Now().Add(et).Unix()

	c := jwt.StandardClaims{
		Issuer:    "xBasis",
		ExpiresAt: expireTime,
	}

	claims := UserClaims{
		Token:          authorize,
		StandardClaims: c,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(tokenKey))

}

func b2s(bs []uint8) string {
	var ba []byte
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

func offlineUser(connectioncli connectioncli.ConnectionClient, repo *tokenRepo, userId, clientId string) *xbasis_commons_dto.State {
	v, err := repo.SizeOf(userId)
	if err != nil {
		return errstate.ErrSystem
	}

	if len(v) > 0 {
		i := len(clientId)
		for _, k := range v {
			key := b2s(k.([]uint8))
			if key[0:i] == clientId {
				offlineTarget(connectioncli, repo, userId, key, clientId)

				//除了需要离线client之外，还需要离线与之相关的relation
				offlineRelation(connectioncli, v, repo, userId, key[i+1:])
			}
		}
	}

	return errstate.Success
}

func offlineRelation(connectioncli connectioncli.ConnectionClient, v []interface{}, repo *tokenRepo, userId, relation string) *xbasis_commons_dto.State {
	for _, k := range v {
		key := b2s(k.([]uint8))
		result := strings.Index(key, ".")
		if key[result+1:] == relation {
			offlineTarget(connectioncli, repo, userId, key, key[0:result])
		}
	}
	return errstate.Success
}

func offlineTarget(connectioncli connectioncli.ConnectionClient, repo *tokenRepo, userId, key, clientId string) {
	repo.Remove(userId, key)
	connectioncli.OfflineToAppClientUser(userId, clientId)
}

package authenticationhandlers

import (
	"github.com/dgrijalva/jwt-go"
	"konekko.me/gosion/authentication/repositories"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"konekko.me/gosion/connection/cmd/connectioncli"
	"strings"
	"time"
)

type UserClaims struct {
	Token *authentication_repositories.SimpleUserToken

	jwt.StandardClaims
}

func decodeToken(token, tokenKey string) (*UserClaims, error) {
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
func encodeToken(tokenKey string, et time.Duration, authorize *authentication_repositories.SimpleUserToken) (string, error) {
	expireTime := time.Now().Add(et).Unix()

	c := jwt.StandardClaims{
		Issuer:    "gosion.authenticate",
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

func sizeCheck(connectioncli connectioncli.ConnectionClient, repo *authentication_repositories.TokenRepo, userId, clientId string) *gs_commons_dto.State {
	v, err := repo.SizeOf(userId)
	if err != nil {
		return errstate.ErrSystem
	}

	if len(v) > 0 {
		i := len(clientId)
		for _, k := range v {
			key := b2s(k.([]uint8))
			if key[0:i] == clientId {
				go offlineTarget(connectioncli, repo, userId, key, clientId)
			}
		}
	}

	return errstate.Success
}

func offlineRelation(connectioncli connectioncli.ConnectionClient, repo *authentication_repositories.TokenRepo, userId, relation string) *gs_commons_dto.State {
	v, err := repo.SizeOf(userId)
	if err != nil {
		return errstate.ErrSystem
	}
	if len(v) > 0 {
		for _, k := range v {
			key := b2s(k.([]uint8))
			result := strings.Index(key, ".")
			if key[result+1:] == relation {
				go offlineTarget(connectioncli, repo, userId, key, key[0:result])
			}
		}
	}
	return errstate.Success
}

func offlineTarget(connectioncli connectioncli.ConnectionClient, repo *authentication_repositories.TokenRepo, userId, key, clientId string) {
	repo.Remove(userId, key)
	connectioncli.OfflineToAppClientUser(userId, clientId)
}

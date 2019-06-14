package oauth

import (
	"fmt"
	"konekko.me/gosion/commons/config/call"
	"konekko.me/gosion/commons/dto"
	"konekko.me/gosion/commons/errstate"
	"net/http"
)

type WechatOAuth struct {
	AppId  string
	Secret string
}

type token struct {
	AccessToken string
	OpenId      string
}

func LoginIntoWechat(code string) *gs_commons_dto.State {
	configuration := serviceconfiguration.Get()
	if configuration.OAuth == nil {
		return errstate.ErrRequest
	}
	oauth := configuration.OAuth
	wechat, ok := oauth["wechat"]
	if !ok {
		return errstate.ErrOAuthTypeNotFound
	}
	info, ok := wechat.(*WechatOAuth)
	if !ok {
		return errstate.ErrRequest
	}
	token, state := getAccessToken(info, code)
}

func GetWechatUnionId() {

}

func getAccessToken(info *WechatOAuth, code string) (string, *gs_commons_dto.State) {
	r, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		info.AppId, info.Secret, code))
	if err != nil {
		return
	}

}

func getOpenId(accessToken string) {

}

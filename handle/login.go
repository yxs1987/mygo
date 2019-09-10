package handle

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
	"mygo/model"
	"mygo/service"
	"mygo/setting"
	"strconv"
	"time"
)

const SecretKey = "abc"

type OpenId struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

func Login(ctx *fasthttp.RequestCtx) {

	var wxresponse model.WechatResponse
	var rs Response

	app_id := setting.APPID
	app_secret := setting.APPSECRET

	url := setting.WECHAT_LOGIN_URL
	err := json.Unmarshal(ctx.Request.Body(), &wxresponse)
	if err != nil {

		fmt.Println("登录出错", err)
	}

	url = fmt.Sprintf(url, app_id, app_secret, wxresponse.Code)
	httpcode, result, err := fasthttp.Get(nil, url)

	if httpcode != fasthttp.StatusOK || err != nil {
		CommonWriteError(ctx, rs)
		return
	}

	jsonStr := string(result)
	//正确的时候返回openid和session_key
	var open OpenId
	jsonerr := json.Unmarshal([]byte(jsonStr), &open)
	if jsonerr != nil {
		CommonWriteError(ctx, rs)
		return
	}

	resp := service.CreateUser(open.OpenId, wxresponse.UserInfo)
	if resp.StatusCode == 200 {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		//有效期7天
		claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
		claims["iat"] = time.Now().Unix()
		claims["iss"] = "yh"
		claims["uid"] = strconv.FormatInt(resp.Data.(int64), 10)

		token.Claims = claims

		tokenString, err := token.SignedString([]byte(SecretKey))

		if err != nil {
			fmt.Println("token错误", err)
		} else {
			rs.Msg = resp.Msg
			rs.StatusCode = resp.StatusCode

			rs.Data = map[string]string{"token": tokenString}
			CommonWriteSuccess(ctx, rs)
		}
	} else {
		CommonWriteError(ctx, rs)
	}

}

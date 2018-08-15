package model

import (
	"github.com/jinzhu/gorm"
	"github.com/bwmarrin/snowflake"
	"log"
	"net/http"
	"io/ioutil"
	"mygo/setting"
	"fmt"
	"encoding/json"

	"crypto/md5"
	"io"
	"math/rand"
)

type WxUser struct {
	Model
	OpenId string `json:"open_id"`
	NickName string `json:"nick_name"`
	AvatarUrl string `json:"avatarUrl"`
	Gender int `json:"gender"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	AddressId int32 `json:"address_id"`
}

//登录返回的结构体
type returnResult struct {
	code string
	msg string
	token string
}

func (user *WxUser) BeforeSave(scope *gorm.Scope) (err error)  {
	node,err := snowflake.NewNode(1)
	if err != nil{
		err = gorm.Errors{}
	}
	id := node.Generate()
	scope.SetColumn("ID",id)
	return nil
}

func WxLogin(data map[string]interface{}) *returnResult {

	st := &returnResult{
		code:"200",
	}

	var code interface{}
	if data["code"] != nil{
		code = data["code"]
	}

	app_id := setting.APPID
	app_secret := setting.APPSECRET

	url := setting.WECHAT_LOGIN_URL

	url = fmt.Sprintf(url, app_id, app_secret, code)
	log.Print(url)
	client := &http.Client{}
	result, err := http.NewRequest("GET", url, nil)
	if err != nil {
		st.code = "400"
		st.msg = err.Error()
	}

	response, _ := client.Do(result)
	body, err := ioutil.ReadAll(response.Body)
	jsonStr := string(body)

	//正确的时候返回openid和session_key
	var f map[string] interface{}
	jsonerr := json.Unmarshal([]byte(jsonStr), &f)
	if jsonerr != nil {
		st.msg = jsonerr.Error()
	}

	session_key := f["session_key"]
	openid := f["openid"]

	var user WxUser
	//创建会员
	findErr := db.Table("rw_user").Model(&user).Where("open_id=?",openid).First(&user).Error;
	if findErr != nil && findErr != gorm.ErrRecordNotFound{
		//报异常错误了
		st.code = "400"
		st.msg = err.Error()
	}

	if err == gorm.ErrRecordNotFound {
		//无会员创建
		if err := db.Table("rw_user").Model(&user).Save(&user).Error;err != nil{
			st.code = "400"
			st.msg = "创建失败"
		}
	} else{
		//加个盐
		salt := fmt.Sprintf("%s_%s_%s",session_key,openid,rand.Float64())
		w := md5.New()
		io.WriteString(w,salt)
		token := fmt.Sprintf("%x", w.Sum(nil))
		//有会员更新token
		st.token = token
		//后续存入redis
	}

	return st
}

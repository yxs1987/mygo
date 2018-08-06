package wx

import (
	"github.com/gin-gonic/gin"
	"mygo/setting"
	"fmt"
	"net/http"
)


func login(c *gin.Context){
	code:= c.DefaultQuery("code","")

	app_id := setting.APPID
	app_secret := setting.APPSECRET

	url:= "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	url = fmt.Sprintf(url,app_id,app_secret,code)
	client := &http.Client{}
	result,err:= client.Get(url)
	if err != nil{
		c.JSON(400,gin.H{
			"msg":err,
		})
	}

	c.JSON(200,result)
}

func aa(){

}
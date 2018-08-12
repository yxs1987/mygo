package wx

import (
	"github.com/gin-gonic/gin"
	"mygo/setting"
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)


func Login(c *gin.Context) {
	code := c.DefaultQuery("code", "")

	app_id := setting.APPID
	app_secret := setting.APPSECRET

	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	url = fmt.Sprintf(url, app_id, app_secret, code)
	log.Print(url)
	client := &http.Client{}
	result, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err,
		})
	}

	response, _ := client.Do(result)

	body, err := ioutil.ReadAll(response.Body)
	jsonStr := string(body)

	var f interface{}
	jsonerr := json.Unmarshal([]byte(jsonStr), &f)
	if jsonerr != nil {
		c.JSON(http.StatusOK, jsonerr)
	}

	c.JSON(http.StatusOK, f)
}
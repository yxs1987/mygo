package wx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mygo/model"
)


func Login(c *gin.Context) {
	//绑定json数据
	var data map[string]interface{}
	c.BindJSON(&data)
	//返回token或者错误信息
	result := model.WxLogin(data)
	c.JSON(http.StatusOK,result)
}
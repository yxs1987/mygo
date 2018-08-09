package router

import (
	"github.com/gin-gonic/gin"
	"mygo/setting"
	"mygo/router/api/wx"
)

func InitRouter() *gin.Engine{

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/",index)

	apiwx := r.Group("/api/wx")
	{
		apiwx.GET("/login",wx.Login)
		apiwx.GET("/good",wx.GoodList)
		apiwx.GET("/good/:id",wx.GoodView)
	}

	return r
}

//小程序基础信息
func index(c *gin.Context){

	c.JSON(200,gin.H{
		"sss":"aa",
	})
}

func JsonSuccess(c *gin.Context,data interface{},msg string){
	c.JSON(200,gin.H{
		"code":200,
		"data":data,
		"msg":msg,
	})
}
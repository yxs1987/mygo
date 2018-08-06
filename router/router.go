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
		apiwx.GET("/login",wx.login)
		apiwx.GET("/goods",)
	}

	return r
}

//小程序基础信息
func index(c *gin.Context){

	c.JSON(200,gin.H{
		"sss":"aa",
	})
}
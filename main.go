package mygo

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/ginS"
)


func main(){
	r := gin.Default()
	ginS.GET("/", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"this is ok",
		})
	})

	r.Run("127.0.0.1:9321")
}
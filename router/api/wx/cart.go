package wx

import (
	"github.com/gin-gonic/gin"
	"log"
	"mygo/model"
)

func AddToCart(c *gin.Context) {

	var maps interface{}
	err:= c.BindJSON(maps);
	if err != nil{
		log.Println(err)
	}

	result := model.AddGoodToCart(maps)
	if result != 0{

	}

	c.JSON(200,gin.H{

		"code":400,
		"msg":"添加成功",
		"data":result,
	})
}
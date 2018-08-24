package wx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mygo/model"
	"github.com/Unknwon/com"
)

func CategoryList(c *gin.Context) {

	data := make(map[string]interface{})

	data["list"] = model.CategoryList();

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": data,
	})
}

func CategoryGoods(c *gin.Context){
	id := com.StrTo(c.Param("id")).MustInt()
	var data interface{}
	data = model.CategoryGoods(id)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": data,
	})
}

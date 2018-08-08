package wx

import (
	"github.com/gin-gonic/gin"
	"mygo/model"
	"net/http"
)

//商品列表
func GoodList(c *gin.Context) {

	name := c.DefaultQuery("name", "")
	category_id := c.DefaultQuery("category_id", "0")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps[name] = name
	}
	if category_id != "0" {
		maps[category_id] = category_id
	}

	data["list"] = model.GoodList(1, 20, maps)
	data["total"], _ = model.Total(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": data,
	})

}

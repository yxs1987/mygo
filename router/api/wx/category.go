package wx

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"mygo/model"
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

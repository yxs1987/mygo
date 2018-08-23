package wx

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"mygo/model"
	"net/http"
	//"log"
	//"log"
	"fmt"
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

func GoodView(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var data interface{}
	data = model.GetDataByPk(id)
	fmt.Println(data)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": data,
	})
}

func GoodView2(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	var data interface{}
	data = model.GetByPk(id)
	fmt.Println(data)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": data,
	})
}

func GoodDel(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	model.DelByPk(id)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "删除",
	})
}

func GoodEdit(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	good := model.Good{GoodsId: id}
	err := c.BindJSON(&good)

	if err != nil {
		bool, err := model.EditByPk(id, good)
		if bool == false {
			c.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "更新成功",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "数据错误",
		})
	}
}

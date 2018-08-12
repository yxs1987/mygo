package wx

import (
	"github.com/gin-gonic/gin"
	"mygo/model"
	"net/http"
	"github.com/Unknwon/com"
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

func GoodView(c *gin.Context){
	id := com.StrTo(c.Param("id")).MustInt()
	var data interface {}
	data = model.GetDataByPk(id)
	fmt.Println(data)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"获取成功",
		"data":data,
	})
}

func GoodEdit(c *gin.Context){
	id := com.StrTo(c.Param("id")).MustInt()

	good := model.Good{GoodsId:id}
	err := c.BindJSON(&good)

	if err != nil{
		bool,err := model.EditByPk(id,good)
		if bool == false{
			c.JSON(http.StatusOK,gin.H{
				"code":400,
				"msg":err,

			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"code":200,
				"msg":"更新成功",

			})
		}
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"msg":"数据错误",
		})
	}
}

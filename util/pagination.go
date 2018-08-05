package util

import (
	"github.com/gin-gonic/gin"
	"github.com/Unknwon/com"
	"mygo/setting"
)

func GetPageNum(c *gin.Context) int {
	result := 0
	pageNum,_ := com.StrTo(c.Query("pageNum")).Int()
	pageSize,_ := com.StrTo(c.Query("pageSize")).Int()

	if pageSize > 0{
		result = (pageNum - 1) * setting.PageSize
	}

	return result
}
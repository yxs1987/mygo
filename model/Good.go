package model

import "github.com/gin-gonic/gin"

type Good struct {
	Model

	GoodsName string
	SpecType int
	CategoryId int
	DeductStockType int
	Content string
	SalesActual int
	SalesInitial int
	GoodsSort int
	DeliveryId int
	Status int
	GoodsStatus int
}

func GoodList(pageNum int,pageSize int,maps interface{}) (goods []Good){
	db.Debug().Where(maps).Offset(pageNum).Limit(pageSize).Find(&goods)
	return
}

func Total(maps interface{}) (int,error)  {
	var count int
	if err := db.Model(&Good{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count,nil
}


//获取单条数据
func GoodView(c *gin.Context) {
	id := c.DefaultQuery("id","0")
	if id == "0" {
		return
	}

	db.Where("id=?",id).Find(&Good{})
	return
}
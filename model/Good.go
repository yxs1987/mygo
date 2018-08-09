package model

import (
	"log"
	"fmt"
)


type Good struct {
	GoodsId int `gorm:"primary_key" json:"goods_id"`
	SpecType int `json:"spec_type"`
	CategoryId int `json:"category_id"`
	DeductStockType int `json:"deduct_stock_type"`
	Content string `json:"content"`
	SalesActual int `json:"sales_actual"`
	SalesInitial int `json:"sales_initial"`
	GoodsSort int `json:"goods_sort"`
	DeliveryId int `json:"delivery_id"`
	Status int `json:"status"`
	GoodsStatus int `json:"goods_status"`
}

func (Good) TableName() string {
	return "rw_goods"
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
func GoodView(id int) (good Good) {
	println(&db)
	if err := db.Debug().
		Where("goods_id=?", id).First(&good).Error;err != nil {
		fmt.Println("ddsdsds")
		log.Fatal(err)
	}else {
		fmt.Println(good)
		fmt.Println("xxx")
	}
	fmt.Println(good)
	return
}


func GoodEdit(id int,data interface{}) bool {
	db.Model(&Good{}).Where("id=?",id).Updates(data)
	return true
}
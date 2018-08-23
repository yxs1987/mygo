package model

import (
	"fmt"
	"log"
)

type Good struct {
	GoodsId         int    `gorm:"primary_key" json:"goods_id"`
	GoodsName       string `json:"goods_name"`
	SpecType        int    `json:"spec_type"`
	CategoryId      int    `json:"category_id"`
	DeductStockType int    `json:"deduct_stock_type"`
	Content         string `json:"content"`
	SalesActual     int    `json:"sales_actual"`
	SalesInitial    int    `json:"sales_initial"`
	GoodsSort       int    `json:"goods_sort"`
	DeliveryId      int    `json:"delivery_id"`
	Status          int    `json:"status"`
	GoodsStatus     int    `json:"goods_status"`
}

func (Good) TableName() string {
	return "rw_goods"
}

func GoodList(pageNum int, pageSize int, maps interface{}) (goods []Good) {
	db.Debug().Table("rw_goods").Where(maps).Offset(pageNum).Limit(pageSize).Find(&goods)
	return
}

func Total(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Good{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

//获取单条数据
func GetDataByPk(id int) (good Good) {
	println(&db)
	if err := db.Debug().
		Where("goods_id=?", id).First(&good).Error; err != nil {
		fmt.Println("ddsdsds")
		log.Fatal(err)
	} else {
		fmt.Println(good)
		fmt.Println("xxx")
	}
	fmt.Println(good)
	return
}

func EditByPk(id int, data interface{}) (bool bool, err error) {
	if err := db.Model(&Good{}).Where("goods_id=?", id).Updates(data).Error; err != nil {
		return false, err
	}
	return true, err
}

func DelByPk(id int) (bool bool, err error) {
	var good Good
	err = db.Model(&good).Where("goods_id=?", id).Error
	if err != nil {
		return false, err
	}
	db.Delete(&good)
	return true, err
}

func GetByPk(id int) (good Good) {
	println(&db)
	if err := db.Debug().
		Where("goods_id=?", id).First(&good).Error; err != nil {
		fmt.Println("ddsdsds")
		log.Fatal(err)
	} else {
		fmt.Println(good)
		fmt.Println("xxx")
	}
	fmt.Println(good)
	return
}

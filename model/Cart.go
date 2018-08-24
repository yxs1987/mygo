package model

import (
	"github.com/goinggo/mapstructure"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type Cart struct{
	GoodId int `json:"good_id"`
	Token string `json:"token"`
	GoodNum int `json:"good_num"`
	GoodSpecId int `json:"good_spec_id"`
}

func AddGoodToCart(maps interface{}) (c *gin.Context){
	var cart Cart
	var user WxUser
	mapstructure.Decode(maps,&cart)
	//查询用户数据是否存在


	err := db.Model(&user).Where("token=?",cart.Token).Find(&user).Error;
	if err != nil && err != gorm.ErrRecordNotFound{
		log.Println(err);
	}

	if err == gorm.ErrRecordNotFound {
		c.JSON(200,gin.H{})
	}
	return
}

//获取购物车中的商品总个数
func GetNum(){

}
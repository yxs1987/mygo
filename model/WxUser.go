package model

import (
	"github.com/jinzhu/gorm"
	"github.com/bwmarrin/snowflake"
)

type WxUser struct {
	Model
	OpenId string `json:"open_id"`
	NickName string `json:"nick_name"`
	AvatarUrl string `json:"avatarUrl"`
	Gender int `json:"gender"`
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
	AddressId int32 `json:"address_id"`
}

func saveUser()  {
	var wxuser = WxUser{}
	db.Create(wxuser)
}

//登录返回的结构体
type result struct {
	code string
	msg string
}

func (user *WxUser) BeforeSave(scope *gorm.Scope) (err error)  {
	node,err := snowflake.NewNode(1)
	if err != nil{
		err = gorm.Errors{}
	}
	id := node.Generate()
	scope.SetColumn("ID",id)
	return nil
}

func WxLogin(code string,user WxUser) *result {

	st := &result{

	}

	//创建会员
	if err := db.Model(&user).Where("open_id=?",user.OpenId).First(&user).Error;err != nil {
		st.code = "400"
		st.msg = err.Error()
	}

	//无会员创建

	if err := db.Model(&user).Save(&user).Error;err != nil{


	}

	//有会员更新token

	return st
}

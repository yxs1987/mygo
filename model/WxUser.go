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

func (user *WxUser) BeforeSave(scope *gorm.Scope) (err error)  {
	node,err := snowflake.NewNode(1)
	if err != nil{
		err = gorm.Errors{}
	}
	id := node.Generate()
	scope.SetColumn("ID",id)
	return nil
}

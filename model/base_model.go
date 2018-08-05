package model

import (
	"github.com/jinzhu/gorm"
	"mygo/setting"
	"log"
	"fmt"
)

var db *gorm.DB


type Model struct{
	ID int `gorm:"primary_key" json:"id"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
}

func init(){
	var (
		err error
		dbType,dbName,dbPassword,dbUser,dbHost,tablePrefix string
	)

	sec,err := setting.Cfg.GetSection("database")

	if err != nil{
		log.Fatal("找不到database模块",err)
	}

	dbType = sec.Key("DB_TYPE").String()
	dbName = sec.Key("DB_NAME").String()
	dbPassword = sec.Key("DB_PASSWORD").String()
	dbUser = sec.Key("DB_USER").String()
	dbHost = sec.Key("DB_HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db,err := gorm.Open(dbType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbName))

	if err!=nil{
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(false)
	db.DB().SetMaxIdleConns(60)
	db.DB().SetMaxOpenConns(60)
}

func closeDB()  {
	defer db.Close()
}
package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	//配置mysql连接参数
	username := "root"
	passwrod := "020903"
	host := "127.0.0.1"
	port := 3306
	Dbname := "gorm"
	dnt := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, passwrod, host, port, Dbname)

	if db, err := gorm.Open(mysql.Open(dnt), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		fmt.Println("mysql  connected err：", err)
	} else {
		DB = db
	}
}

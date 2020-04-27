package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"tasker/server/model"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "192.168.1.100"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database , err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

// 获取连接实例
func GetDB() *gorm.DB {
	return DB
}

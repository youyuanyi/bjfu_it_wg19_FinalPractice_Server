package common

import (
	"WeatherServer/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/url"
)

/* common/database.go */
var DB *gorm.DB

func InitDB() *gorm.DB {
	driverName := "mysql"
	user := "root"
	password := "YJC706989"
	host := "1.15.56.246"
	port := "3306"
	database := "ZhuanYeShiJian"
	charset := "utf8"
	loc := "Asia/Shanghai"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		user,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	// 连接数据库
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to open database: " + err.Error())
	}
	// 迁移数据表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Userequipments{})
	db.AutoMigrate(&model.Area{})
	db.AutoMigrate(&model.Node{})
	db.AutoMigrate(&model.Data{})
	db.AutoMigrate(&model.Physical{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

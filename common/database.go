package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kok/model"
)

var DB *gorm.DB

// 数据库连接池
func InitDB() *gorm.DB {
	driverName := "mysql"   //数据库类型
	host := "localhost"     //数据库地址
	port := "3306"          //数据库端口
	database := "kok"       //数据库名称
	username := "root"      //数据库用户名
	password := "lomi1015." //数据库密码
	charset := "utf8"       //数据库编码方式
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.New(mysql.Config{DriverName: driverName, DSN: args}), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

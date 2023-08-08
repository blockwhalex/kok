package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kok/model"
)

var DB *gorm.DB

// 数据库连接池
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName") //数据库类型
	host := viper.GetString("datasource.host")             //数据库地址
	port := viper.GetString("datasource.port")             //数据库端口
	database := viper.GetString("datasource.database")     //数据库名称
	username := viper.GetString("datasource.username")     //数据库用户名
	password := viper.GetString("datasource.password")     //数据库密码
	charset := viper.GetString("datasource.charset")       //数据库编码方式
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

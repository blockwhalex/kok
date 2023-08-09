package model

import "gorm.io/gorm"

// 用户信息model
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`             //用户名
	UUID      string `gorm:"type:varchar(255);not null;primaryKey"` //唯一标识
	Email     string `gorm:"type:varchar(50);not null;unique"`      //邮箱
	Telephone string `gorm:"type:varchar(11);not null;unique"`      //手机号
	Password  string `gorm:"type:varchar(255);not null"`            //密码
}

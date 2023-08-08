package model

import "gorm.io/gorm"

// 用户信息model
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`        //用户名
	Telephone string `gorm:"type:varchar(11);not null;unique"` //手机号
	Password  string `gorm:"type:varchar(255);not null"`       //密码
}

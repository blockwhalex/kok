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
	State     int    `gorm:"type:DEFAULT 0;not null"`               //状态：0-正常，1-失效，2-禁用
	Level     int    `gorm:"type:DEFAULT 0;not null"`               //等级：0-无，1-一级，2-二级
	Integral  int    `gorm:"type:DEFAULT 0;"`                       //积分

}

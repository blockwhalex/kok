package dto

import "kok/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

// 控制数据传输对象的展示
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

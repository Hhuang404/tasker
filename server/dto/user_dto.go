package dto

import "tasker/server/model"

type UserDto struct {
	Name     string `json:"name"`
	Nickname string `json:"nickName"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:     user.Username,
		Nickname: user.Nickname,
	}
}

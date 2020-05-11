package model

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Username  string `json:"userName"`
	Password  string `json:"password"`
	Nickname  string `json:"nickName"`
}

type UserToken struct {
	Name     string `json:"name"`
	Nickname string `json:"nickName"`
}

func ToUserDto(user User) UserToken {
	return UserToken{
		Name:     user.Username,
		Nickname: user.Nickname,
	}
}

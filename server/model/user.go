package model

import "time"

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Username  string `json:"userName"`
	Password  string `json:"password"`
	Nickname  string `json:"nickName"`
}

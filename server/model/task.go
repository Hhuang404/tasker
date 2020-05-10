package model

import (
	"time"
)

type Task struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

// 列表
type TaskListVo struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
}

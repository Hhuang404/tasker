package model

import (
	"time"
)

// task
type Task struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	UserId    uint      `json:"userId"`
}

// task 明细
type TaskDetail struct {
	ID         uint      `json:"id"`
	Status     uint8     `json:"status"`
	TaskId     uint      `json:"taskId"`
	UserId     uint      `json:"userId"`
	Name       string    `json:"name"`
	Info       string    `json:"info"`
	UpdateAt   time.Time `json:"updateAt"`
	CreatedAt  time.Time `json:"createdAt"`
	DeadlineAt time.Time `json:"deadlineAt"`
}

// 列表
type TaskListVo struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	Name      string `json:"name"`
}

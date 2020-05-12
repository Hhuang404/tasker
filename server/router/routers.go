package router

import (
	"github.com/gin-gonic/gin"
	"tasker/server/api"
	"tasker/server/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 用户
	r.POST("/api/auth/register", api.Register) // 用户注册
	r.POST("/api/auth/login", api.Login)       // 用户登录
	//r.GET("/api/auth/info", middleware.AuthMiddleware(), api.Info) //获取当前用户信息

	// task
	r.GET("/api/task/:userId", middleware.AuthMiddleware(), api.GetTaskList) // 获取 task 列表
	r.POST("/api/task/add", middleware.AuthMiddleware(), api.AddTask)        // 添加 task

	// task_detail
	r.GET("/api/task/detail/:taskId", middleware.AuthMiddleware(), api.GetTaskDetailList) // 获取 task 详情列表


	return r
}

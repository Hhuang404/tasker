package api

import (
	"github.com/gin-gonic/gin"
	"tasker/server/common"
	"tasker/server/constant"
	"tasker/server/model"
	"tasker/server/response"
	"tasker/server/utils"
)

// 获取 task detail
func GetTaskDetailList(ctx *gin.Context) {
	var db = common.DB
	taskId := ctx.Param("taskId")
	if utils.ParamIsNull(taskId) {
		response.Fail("taskId 为空", ctx)
	}
	var taskDetail []*model.TaskDetail
	db.Where("task_id = ?", taskId).Find(&taskDetail)
	response.Success(constant.QuerySuccess, taskDetail, ctx)
}

//// 添加 task detail
//func AddTaskDetail(ctx *gin.Context){
//
//
//}

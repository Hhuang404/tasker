package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tasker/server/common"
	"tasker/server/constant"
	"tasker/server/model"
	"tasker/server/response"
	"tasker/server/utils"
)

const (
	// 1 2 3 4 5 6 7 记忆
	YYYYmmDDHHssMM = "2006-01-02 15:04:05"
)

func GetList(ctx *gin.Context) {
	var db = common.DB
	uid := ctx.Param("userId")
	if utils.ParamIsNull(uid) {
		response.Fail("userId 为空", ctx)
	}
	var tasks []*model.Task
	db.Where("user_id = ?", uid).Find(&tasks)
	taskListVo := make([]*model.TaskListVo, 0)
	for _, task := range tasks {
		timeStr := task.CreatedAt.Format(YYYYmmDDHHssMM)
		fmt.Println("时间： ", timeStr)
		taskVo := &model.TaskListVo{
			CreatedAt: timeStr,
			ID:        task.ID,
			Name:      task.Name,
		}
		taskListVo = append(taskListVo, taskVo)
	}
	response.Success(constant.QuerySuccess, taskListVo, ctx)
}

package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"tasker/server/common"
	"tasker/server/constant"
	"tasker/server/model"
	"tasker/server/response"
	"tasker/server/utils"
)

// 获取任务列表
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
		timeStr := task.CreatedAt.Format(constant.YearMonthDayHourMinuteSecond)
		taskVo := &model.TaskListVo{
			CreatedAt: timeStr,
			ID:        task.ID,
			Name:      task.Name,
		}
		taskListVo = append(taskListVo, taskVo)
	}
	response.Success(constant.QuerySuccess, taskListVo, ctx)
}

// 添加任务
func AddTask(ctx *gin.Context) {
	var db = common.DB
	name := ctx.PostForm("name")
	userId := ctx.PostForm("userId")
	if utils.ParamIsNull(name, userId) {
		response.Fail(constant.ParamNull, ctx)
		return
	}
	uid, err := strconv.Atoi(userId)
	if err != nil {
		response.Fail(constant.ParamError, ctx)
		return
	}
	if db.Create(&model.Task{
		Name:   name,
		UserId: uint(uid),
	}).Error != nil {
		response.Fail(constant.DuplicateEntity, ctx)
		return
	}
	response.Success(constant.AddSuccess, nil, ctx)
}

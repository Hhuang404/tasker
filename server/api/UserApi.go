package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"tasker/server/common"
	"tasker/server/constant"
	"tasker/server/dto"
	"tasker/server/model"
	"tasker/server/response"
	"tasker/server/utils"
)

// 用户注册
func Register(ctx *gin.Context) {
	var DB = common.GetDB()
	// 获取参数
	username := ctx.PostForm("username")
	nickname := ctx.PostForm("nickname")
	password := ctx.PostForm("password")
	if utils.ParamIsNull(username, nickname, password) {
		response.Fail("username , password 或 nickname 为空", ctx)
		return
	}
	if isUsernameExist(DB, username) {
		response.Fail("用户已存在", ctx)
		return
	}
	// create user
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Fail("加密错误", ctx)
		return
	}
	newUser := model.User{
		Username: username,
		Nickname: nickname,
		Password: string(hasPassword),
	}
	DB.Create(&newUser)
	response.Success("注册成功", nil, ctx)
}

// 用户登录
func Login(ctx *gin.Context) {
	var DB = common.GetDB()
	// 获取参数
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// 判断手机号是否存在
	var user model.User
	DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		response.Fail("用户不存在", ctx)
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail("密码错误", ctx)
		return
	}
	// 密码正确发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Fail("系统错误", ctx)
		log.Printf("token generate error : %v", err)
	}
	ctx.SetCookie("token", "Bearer "+token, constant.MaxAgeSecond, "/", constant.DoMain, false, true)
	// 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":  1,
		"msg":   "登录成功",
		"token": token,
	})
}

// 获取当前用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"username": dto.ToUserDto(user.(model.User))}})
}

// 判断登录账号是否存在
func isUsernameExist(db *gorm.DB, username string) bool {
	var user model.User
	db.Where("username = ?", username).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

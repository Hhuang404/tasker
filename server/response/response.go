package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//{
//	code:20001,
//	data:xxxx
//	msg:xx
//}
// 自定义返回
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

// 成功返回
func Success(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// 失败返回
func Fail(ctx *gin.Context, data gin.H, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}

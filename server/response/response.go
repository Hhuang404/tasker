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

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 1
	ERROR   = 0
)

// 自定义返回
func result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 成功返回
func Success(msg string, data interface{}, ctx *gin.Context) {
	result(SUCCESS, data, msg, ctx)
}

// 失败返回
func Fail(msg string, ctx *gin.Context) {
	result(ERROR, nil, msg, ctx)
}

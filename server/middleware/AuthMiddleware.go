package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"tasker/server/common"
	"tasker/server/constant"
	"tasker/server/model"
)

// jwt 验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("token")
		}

		if !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": constant.Unauthorized})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": constant.Unauthorized})
			ctx.Abort()
			return
		}

		// 验证通过验证后获取claim中的 userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": constant.Unauthorized})
			ctx.Abort()
			return
		}

		// 用户存在 将 user 的信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}

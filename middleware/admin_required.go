package middleware

import (
	"github.com/gin-gonic/gin"
	"hulujia/service"
)

// 添加后台管理员中间件，避免用户token误传
func AdminRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := service.UserService.GetCurrentUser(ctx)
		if user == nil {
			ctx.AbortWithStatusJSON(401, gin.H{
				"code":   	401,
				"message": "管理员信息不存在",
			})
			return
		}
		// 处理权限
		ctx.Next()
	}
}

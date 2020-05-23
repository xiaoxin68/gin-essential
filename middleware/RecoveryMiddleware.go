package middleware

import (
	"gin-essential/response"
	"github.com/gin-gonic/gin"
)

//拦截处理panic
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, gin.H{"err": err}, "系统故障")
			}
		}()
		ctx.Next()
	}
}

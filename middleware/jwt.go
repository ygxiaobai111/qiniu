package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"www.github.com/ygxiaobai111/qiniu/pkg/e"
	"www.github.com/ygxiaobai111/qiniu/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS {
			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

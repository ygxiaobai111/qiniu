package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	e2 "www.github.com/ygxiaobai111/qiniu/server/pkg/e"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
)

// 身份认证
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		code = e2.SUCCESS
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)

			if err != nil {
				code = e2.ErrorAuthToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e2.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e2.SUCCESS {

			ctx.JSON(200, gin.H{
				"status": code,
				"msg":    e2.GetMsg(code),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

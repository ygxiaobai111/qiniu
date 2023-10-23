package routes

import (
	"net/http"
	"www.github.com/ygxiaobai111/qiniu/api"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	//加载静态页面
	r.LoadHTMLGlob("view/*")
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"code": 200,
			})
		})

		uG := r.Group("user")

		uG.POST("register", api.UserRegister)

	}
	return r
}

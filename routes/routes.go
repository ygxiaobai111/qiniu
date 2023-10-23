package routes

import (
	"net/http"
	"www.github.com/ygxiaobai111/qiniu/api"
	"www.github.com/ygxiaobai111/qiniu/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	//允许跨域请求
	r.Use(middleware.Cors())
	//加载静态页面
	r.LoadHTMLGlob("view/*")
	//静态资源服务
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"code": 200,
			})
		})
		//创建user组
		uG := r.Group("user")
		//用户注册接口
		uG.POST("register", api.UserRegister)

	}
	return r
}

package routes

import (
	"net/http"
	"www.github.com/ygxiaobai111/qiniu/server/api"
	"www.github.com/ygxiaobai111/qiniu/server/middleware"

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
		{
			uG.GET("info/:id")
			//注册
			uG.POST("register", api.UserRegister)
			//登录
			uG.POST("login")
			//关注
			uG.POST("action")
			//关注列表
			uG.GET("follow/list/:id")
			//粉丝列表
			uG.GET("follower/list/:id")
			//好友列表
			uG.GET("friend/list/:id")

		}
		vG := v1.Group("video")
		{
			vG.GET("feed")
			//用户视频列表
			vG.GET("publish/list/:id")
			//用户投稿
			vG.POST("publish/action")
		}
		//收藏夹
		v1.GET("favlist/:uid/:fid")
		//点赞
		v1.POST("favorite/action/")
		//用户喜欢列表
		v1.GET("/favorite/list/:id")
		//评论
		v1.POST("comment/action/")
		//评论列表
		v1.GET("comment/list/:id")
		//弹幕发送
		v1.POST("barrage")
	}
	return r
}

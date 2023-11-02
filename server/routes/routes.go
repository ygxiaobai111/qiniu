package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"www.github.com/ygxiaobai111/qiniu/server/api"
	"www.github.com/ygxiaobai111/qiniu/server/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	//允许跨域请求
	r.Use(middleware.Cors())
	//加载静态页面
	//r.LoadHTMLGlob("view/*")
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
			//用户信息
			uG.GET("info", api.UserInfo)
			//注册
			uG.POST("register", api.UserRegister)
			//登录
			uG.POST("login", api.UserLogin)
			//关注/取关
			uG.POST("action", api.UserAction)
			//关注列表
			uG.GET("follow/list", api.UserFollow)
			//粉丝列表
			uG.GET("follower/list", api.UserFollower)
			//好友列表
			uG.GET("friend/list", api.UserFriend)

		}
		vG := v1.Group("video")
		{
			//搜索
			vG.GET("search", api.VideoSearch)
			//视频分类
			vG.GET("channel/:id", api.VideoChannel)
			//热门视频
			vG.GET("hot", api.VideoHot)
			//视频流
			vG.GET("feed", api.VideoFeed)
			//用户视频列表
			vG.GET("publish/list", api.VideoGetPublish)
			//用户投稿
			vG.POST("publish/action", api.VideoCreate)
			vG.PUT("publish/action", api.VideoUpdatePublish)
			vG.DELETE("publish/action", api.VideoDelPublish)
			//历史视频
			vG.GET("before", api.VideoBefore)

		}
		iG := v1.Group("interaction")
		{
			//收藏夹
			iG.GET("favlist", api.GetFavlist)

			//用户喜欢列表
			iG.GET("favorite", api.GetFavorite)

			//评论列表
			iG.GET("comment", api.GetComment)

			//弹幕获取
			iG.GET("barrage", api.GetBarrage)

			iG.Use()
			//创建收藏夹
			iG.POST("favlist", api.FavlistCreate)
			//加入收藏夹
			iG.PUT("fav", api.FavlistAdd)
			//退出收藏夹
			iG.DELETE("fav", api.FavlistDel)
			//删除收藏夹
			iG.DELETE("favlist", api.DelFavlist)
			//评论
			iG.POST("comment", api.CommentCreate)
			//点赞/取消点赞
			iG.POST("favorite", api.Favorite)
			//弹幕发送
			iG.POST("barrage", api.Barrage)
			//用户画像
			iG.POST("personas", api.Personas)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/service"
	"www.github.com/ygxiaobai111/qiniu/server/types"
)

// @Summary 创建用户
// @Description 通过表单提交创建用户
// @ID VideoCreate
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} UserRegisterResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/register [post]

func VideoCreate(ctx *gin.Context) {
	var req types.VideoCreateReq
	if err := ctx.ShouldBind(&req); err != nil {
		// 参数校验
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	// 获取userSrv对象
	srv := service.GetVideoSrv()
	form, _ := ctx.MultipartForm()
	images := form.File["image"]

	videos := form.File["video"]
	if len(images) > 1 && len(videos) != 1 {
		err := os.ErrInvalid
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}

	resp, err := srv.VideoCreate(ctx.Request.Context(), videos[0], images[0])
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp))
}

// @Summary 视频搜索
// @Description 通过表单提交进行视频搜索
// @ID VideoSearch
// @Accept x-www-form-urlencoded
// @Produce json
// @Param text formData int64 true "关键字"
// @Header 200 {string} Token  "我的token"
// @Success 200 {object} VideoSearchResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/search [get]

type VideoSearchResponse types.Response

func VideoSearch(ctx *gin.Context) {
	var req *types.VideoSearch
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetVideoSrv()

	resp, err := srv.VideoSearch(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 视频分类
// @Description 通过表单提交获取该分类的视频
// @ID VideoChannel
// @Accept x-www-form-urlencoded
// @Produce json
// @Param channel_id formData int64 true "分类id"
// @Header 200 {string} Token "我的token"
// @Success 200 {object} VideoChannelResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/channel [post]

type VideoChannelResponse types.Response

func VideoChannel(ctx *gin.Context) {
	var req *types.VideoChannel
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetVideoSrv()

	resp, err := srv.VideoChannel(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 用户视频列表
// @Description 通过表单获取该用户发布的视频
// @ID VideoPublish
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id formData int64 true "用户id"
// @Header 200 {string} Token true "我的token"
// @Success 200 {object} VideoPublishResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/Publish [post]

type VideoGetPublishResponse types.Response

func VideoGetPublish(ctx *gin.Context) {
	var req *types.VideoGetPublish
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetVideoSrv()

	resp, err := srv.VideoGetPublish(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 用户视频更新
// @Description 通过表单提交用户发布的视频
// @ID VideoPublish
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id formData int64 true "用户id"
// @Header 200 {string} Token true "我的token"
// @Success 200 {object} VideoUpdatePublishResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/Publish [post]

type VideoUpdatePublishResponse types.Response

func VideoUpdatePublish(ctx *gin.Context) {
	var req *types.VideoUpdatePublish
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetVideoSrv()

	resp, err := srv.VideoUpdatePublish(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 删除用户视频
// @Description 通过表单删除用户的视频
// @ID VideoDelPublish
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id formData int64 true "用户id"
// @Header 200 {string} Token true "我的token"
// @Success 200 {object} VideoDelPublishResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/Publish [post]

type VideoDelPublishResponse types.Response

func VideoDelPublish(ctx *gin.Context) {
	var req *types.VideoDelPublish
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetVideoSrv()

	resp, err := srv.VideoDelPublish(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 用户历史视频
// @Description 通过表单获取该用户的历史视频
// @ID VideoBefore
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id formData int64 true "用户id"
// @Header 200 {string} Token true "我的token"
// @Success 200 {object} VideoBeforeResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/Before [post]

type VideoBeforeResponse types.Response

func VideoBefore(ctx *gin.Context) {
	var req *types.VideoBefore
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetVideoSrv()

	resp, err := srv.VideoBefore(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

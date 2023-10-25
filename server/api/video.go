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
// @ID UserRegister
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
// @ID UserAction
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
// @ID UserAction
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id formData int64 true "分类id"
// @Header 200 {string} Token "我的token"
// @Success 200 {object} VideoChannelResponse
// @Failure 400 {object} ErrorResponse
// @Router /video/channel [post]

type VideoChannelResponse types.Response

func VideoChannel(ctx *gin.Context) {
	var req *types.UserLoginReq
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

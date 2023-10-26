package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/service"
	"www.github.com/ygxiaobai111/qiniu/server/types"
)

// @Summary 查看收藏夹
// @Description 通过表单提交方式查看收藏夹
// @ID UserRegister
// @Accept x-www-form-urlencoded
// @Produce json
// @Param UserId formData int true "用户id"
// @Param FavlistId formData int  "收藏夹id 不带该参数则返回所有收藏夹"
// @Success 200 {object} GetFavlistResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func GetFavlist(ctx *gin.Context) {
	var req *types.GetFavlistReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 喜欢列表
// @Description 通过表单提交方式查看喜欢列表
// @ID GetFavorite
// @Accept x-www-form-urlencoded
// @Produce json
// @Param UserId formData int true "用户id"
// @Success 200 {object} GetFavoriteResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func GetFavorite(ctx *gin.Context) {
	var req *types.GetFavoriteReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 视频评论列表
// @Description 通过表单提交查看评论列表
// @ID GetComment
// @Accept x-www-form-urlencoded
// @Produce json
// @Param VideoId formData int true "视频id"
// @Success 200 {object} GetCommentResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func GetComment(ctx *gin.Context) {
	var req *types.GetCommentReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 弹幕获取
// @Description 通过表单提交获取弹幕
// @ID GetBarrage
// @Accept x-www-form-urlencoded
// @Produce json
// @Param VideoId formData int true "视频id"
// @Success 200 {object} GetBarrageResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func GetBarrage(ctx *gin.Context) {
	var req *types.GetBarrageReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 创建收藏夹
// @Description 通过表单提交创建收藏夹
// @ID UserRegister
// @Accept x-www-form-urlencoded
// @Produce json
// @Param FavlistName formData string true "收藏夹名称"
// @Param Type formData int true "收藏夹类型，1为公开 2为隐藏"
// @Success 200 {object} FavlistCreateResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func FavlistCreate(ctx *gin.Context) {
	var req *types.FavlisCreatetReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 加入收藏夹
// @Description 通过表单提交将视频加入收藏夹
// @ID UserRegister
// @Accept x-www-form-urlencoded
// @Produce json
// @Param FavlistId formData int true "收藏夹id"
// @Param VideoId formData int true "视频id"
// @Success 200 {object} FavlistAddResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func FavlistAdd(ctx *gin.Context) {
	var req *types.FavlistAddReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 视频移除收藏夹
// @Description 通过表单提交将视频移除收藏夹
// @ID FavlistDel
// @Accept x-www-form-urlencoded
// @Produce json
// @Param FavlistId formData string true "收藏夹id"
// @Param VideoId formData string true "视频id"
// @Success 200 {object} FavlistDelResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func FavlistDel(ctx *gin.Context) {
	var req *types.FavlistDelReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 删除收藏夹
// @Description 通过表单提交删除收藏夹
// @ID DelFavlist
// @Accept x-www-form-urlencoded
// @Produce json
// @Param FavlistId formData string true "收藏夹id"
// @Success 200 {object} DelFavlistResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func DelFavlist(ctx *gin.Context) {
	var req *types.DelFavlistReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 创建评论
// @Description 通过表单提交创建评论
// @ID CommentCreateReq
// @Accept x-www-form-urlencoded
// @Produce json
// @Param VideoId formData int true "视频id"
// @Param Content formData string true "内容"
// @Success 200 {object} CommentCreateResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func CommentCreate(ctx *gin.Context) {
	var req *types.CommentCreateReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 创建用户
// @Description 通过表单提交创建用户
// @ID UserRegister
// @Accept x-www-form-urlencoded
// @Produce json
// @Param VideoId formData int true "视频id"
// @Param Type formData int true "1点赞 2取消"
// @Success 200 {object} FavoriteResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func Favorite(ctx *gin.Context) {
	var req *types.FavoriteReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 弹幕发送
// @Description 通过表单提交弹幕
// @ID UserRegister
// @Accept x-www-form-urlencoded
// @Produce json
// @Param VideoID formData int true " 弹幕所属视频的ID"
// @Param Content formData string true "弹幕内容"
// @Param Color formData string true "弹幕颜色"
// @Param Timestamp formData int true "弹幕出现的时间戳"
// @Success 200 {object} BarrageResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func Barrage(ctx *gin.Context) {
	var req *types.BarrageReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www.github.com/ygxiaobai111/qiniu/server/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/server/service"
	"www.github.com/ygxiaobai111/qiniu/server/types"
)

// ShowAccount godoc
//
//	@Summary		查看收藏夹
//	@Description	查看收藏夹
//	@Tags			交互
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		int	true	"用户id"
//	@Param			favlist_id	query		int	false	"收藏夹id"
//	@Success		200	{object}	types.GetFavlistResp
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/interaction/favlist [get]
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
	srv := service.GetInterSrv()

	resp, err := srv.GetFavlist(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		喜欢列表
// @Description	查看喜欢列表
// @ID				GetFavorite
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			user_id	query		int	true	"用户id"
// @Success		200		{object}	types.GetFavResp
// @Failure		400		{object}	ErrorResponse
// @Router			/interaction/favorite [get]
func GetFavorite(ctx *gin.Context) {
	var req *types.GetFavoriteReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetInterSrv()

	resp, err := srv.GetFavorite(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		视频评论列表
// @Description	提交查看评论列表
// @ID				GetComment
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			video_id	query		int	true	"视频id"
// @Success		200		{object}	types.GetCommentResp
// @Failure		400		{object}	ErrorResponse
// @Router			/interaction/comment [get]
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
	srv := service.GetInterSrv()

	resp, err := srv.GetComment(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		弹幕获取
// @Description	获取弹幕
// @ID				GetBarrage
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			video_id	query		int	true	"视频id"
// @Success		200		{object}	types.GetBarrageResp
// @Failure		400		{object}	ErrorResponse
// @Router			/interaction/barrage [get]
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
	srv := service.GetInterSrv()

	resp, err := srv.GetBarrage(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		创建收藏夹
// @Description	提交创建收藏夹
// @ID				FavlistCreate
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			favlist_name	formData	string	true	"收藏夹名称"
// @Param			type		formData	int		true	"收藏夹类型，1为公开 2为隐藏"
// @Header			200			{string}	Token	"我的token"
// @Success		200			{object}	Response
// @Failure		400			{object}	ErrorResponse
// @Router			/interaction/favlist [post]
func FavlistCreate(ctx *gin.Context) {
	var req *types.FavlisCreatetReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}

	srv := service.GetInterSrv()

	resp, err := srv.FavlistCreate(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		加入收藏夹
// @Description	提交将视频加入收藏夹
// @ID				FavlistAdd
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			request	body		types.FavlistAddReq	true	"想要添加的视频id和文件夹id"
// @Header			200		{string}	Token				"我的token"
// @Success		200		{object}	Response
// @Failure		400		{object}	ErrorResponse
// @Router			/interaction/fav [put]
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
	srv := service.GetInterSrv()

	resp, err := srv.FavlistAdd(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		视频移除收藏夹
// @Description	提交将视频移除收藏夹
// @ID				FavlistDel
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			favlist_id	formData	string	true	"收藏夹id"
// @Param			video_id		formData	string	true	"视频id"
// @Header			200			{string}	Token	"我的token"
// @Success		200			{object}	Response
// @Failure		400			{object}	ErrorResponse
// @Router			/interaction/fav [delete]
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
	srv := service.GetInterSrv()

	resp, err := srv.FavlistDel(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		删除收藏夹
// @Description	提交删除收藏夹
// @ID				DelFavlist
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			favlist_id	formData	string	true	"收藏夹id"
// @Header			200			{string}	Token	"我的token"
// @Success		200			{object}	Response
// @Failure		400			{object}	ErrorResponse
// @Router			/interaction/favlist [delete]
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
	srv := service.GetInterSrv()

	resp, err := srv.DelFavlist(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		创建评论
// @Description	提交创建评论
// @ID				CommentCreateReq
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			video_id	formData	int		true	"视频id"
// @Param			content	formData	string	true	"内容"
// @Header			200		{string}	Token	"我的token"
// @Success		200		{object}	Response
// @Failure		400		{object}	ErrorResponse
// @Router			/interaction/comment [post]
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
	srv := service.GetInterSrv()

	resp, err := srv.CommentCreate(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		点赞/取消赞
// @Description	提交点赞/取消赞
// @ID				Favorite
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			video_id	formData	int		true	"视频id"
// @Param			type	formData	int		true	"1点赞 2取消"
// @Header			200		{string}	Token	"我的token"
// @Success		200		{object}	Response
// @Failure		400		{object}	ErrorResponse
// @Router			/interaction/favorite [post]
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
	srv := service.GetInterSrv()

	resp, err := srv.Favorite(ctx.Request.Context(), req, util.GetUidInToken(ctx))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		弹幕发送
// @Description	提交弹幕
// @ID				Barrage
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			video_id		formData	int		true	" 弹幕所属视频的ID"
// @Param			content		formData	string	true	"弹幕内容"
// @Param			color		formData	string	true	"弹幕颜色"
// @Param			timestamp	formData	int		true	"弹幕出现的时间戳"
// @Header			200			{string}	Token	"我的token"
// @Success		200			{object}	Response
// @Failure		400			{object}	ErrorResponse
// @Router			/interaction/barrage [post]
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
	srv := service.GetInterSrv()

	resp, err := srv.Barrage(ctx.Request.Context(), req, uint64(util.GetUidInToken(ctx)))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary		用户喜好上传
// @Description	提交用户感兴趣视频标签
// @ID				Personas
// @Accept			x-www-form-urlencoded
// @Produce		json
// @Param			category_id	formData	int		true	"标签id"
// @Header			200			{string}	Token	"我的token"
// @Success		200			{object}	Response
// @Failure		400			{object}	ErrorResponse
// @Router			/interaction/personas [post]
func Personas(ctx *gin.Context) {
	var req *types.PersonasReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	// 获取userSrv对象
	srv := service.GetInterSrv()

	resp, err := srv.Personas(ctx.Request.Context(), req, uint64(util.GetUidInToken(ctx)))
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

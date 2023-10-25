package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
// @Router /user/register [post]

type UserRegisterResponse types.Response
type ErrorResponse types.Response

func UserRegister(ctx *gin.Context) {
	var req *types.UserRegisterReq
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

// @Summary 用户登录
// @Description 通过表单提交进行用户登录
// @ID UserLogin
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} UserLoginResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/login [post]

type UserLoginResponse types.Response

func UserLogin(ctx *gin.Context) {
	var req *types.UserLoginReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserLogin(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 用户信息
// @Description 通过对方id和我的token获取对方用户信息
// @ID UserInfo
// @Accept json
// @Produce json
// @Param id query int true "对方用户ID"
// @Header 200 {string} Token "我的token"
// @Success 200 {object} UserInfoResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/login [post]

type UserInfoResponse types.Response

func UserInfo(ctx *gin.Context) {
	var req *types.UserLoginReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserLogin(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 用户关注/取关
// @Description 通过表单提交进行关注/取关
// @ID UserAction
// @Accept x-www-form-urlencoded
// @Produce json
// @Param user_id formData int64 true "对方用户id"
// @Header 200 {string} Token true "我的token"
// @Success 200 {object} UserActionResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/action [post]

type UserActionResponse types.Response

func UserAction(ctx *gin.Context) {
	var req *types.UserLoginReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserLogin(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 关注列表
// @Description 通过userId查询用户关注列表
// @ID UserFollow
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id query int true "用户ID"
// @Header 200 {string} Token "我的token"
// @Success 200 {object} UserFollowResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/follow [get]

type UserFollowResponse types.Response

func UserFollow(ctx *gin.Context) {
	var req *types.UserLoginReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserLogin(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 粉丝列表
// @Description 通过userId查询用户粉丝列表
// @ID UserFollower
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id query int true "用户ID"
// @Header 200 {string} Token "我的token"
// @Success 200 {object} UserFollowerResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/follower [get]

type UserFollowerResponse types.Response

func UserFollower(ctx *gin.Context) {
	var req *types.UserLoginReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserLogin(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

// @Summary 用户好友列表
// @Description 通过用户id查询用户粉丝列表
// @ID UserFriend
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id query int true "用户ID"
// @Header 200 {string} Token "我的token"
// @Success 200 {object} UserFriendResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/friend [get]

type UserFriendResponse types.Response

func UserFriend(ctx *gin.Context) {
	var req *types.UserLoginReq
	//ctx.ShouldBind(&req) 获取前端输入的表单信息
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, types.ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return
	}
	// 获取userSrv对象
	srv := service.GetUserSrv()

	resp, err := srv.UserLogin(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Error(err)
		ctx.JSON(http.StatusOK, types.ErrorResponse(err))
		return
	}
	//返回给前端相应信息
	ctx.JSON(http.StatusOK, types.RespSuccess(ctx, resp, http.StatusOK))
}

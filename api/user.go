package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www.github.com/ygxiaobai111/qiniu/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/service"
	"www.github.com/ygxiaobai111/qiniu/types"
)

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

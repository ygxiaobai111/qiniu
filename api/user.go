package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"www.github.com/ygxiaobai111/qiniu/pkg/util"
	"www.github.com/ygxiaobai111/qiniu/serializer"
	"www.github.com/ygxiaobai111/qiniu/service"
	"www.github.com/ygxiaobai111/qiniu/types"
)

func UserRegister(ctx *gin.Context) {
	var req *types.UserRegisterReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		//打日志
		util.LogrusObj.Error(err)
		return

	}
	srv := service.GetUserSrv()

	resp, err := srv.UserRegister(ctx.Request.Context(), req)
	if err != nil {
		util.LogrusObj.Infoln(err)
		ctx.JSON(http.StatusOK, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.RespSuccess(ctx, resp, 200))
}

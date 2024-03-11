package api

import (
	"ChargPiles/pkg/e"
	"ChargPiles/pkg/utils/ctl"
	"ChargPiles/service"
	"ChargPiles/types"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		srv := service.GetUserSrv()
		resp, err := srv.UserRegister(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}

}

func UserVerificationCodeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserVerificationCodeReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}

		srv := service.GetUserSrv()
		srv.UserVerificationCode(ctx.Request.Context(), &req)
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, nil))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		srv := service.GetUserSrv()
		resp, err := srv.UserLogin(ctx.Request.Context(), &req)
		if err != nil {
			ctx.JSON(http.StatusOK, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func UserUpdateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoUpdateReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			//log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		srv := service.GetUserSrv()
		resp, err := srv.UserInfoUpdate(ctx.Request.Context(), &req)
		if err != nil {
			//log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func ShowUserInfoHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserInfoShowReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			//log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		srv := service.GetUserSrv()
		resp, err := srv.UserInfoShow(ctx.Request.Context(), &req)
		if err != nil {
			//log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

func UploadAvatarHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			//log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}
		file, fileHeader, _ := ctx.Request.FormFile("file")
		if fileHeader == nil {
			err := errors.New(e.GetMsg(e.ErrorUploadFile))
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			//log.LogrusObj.Infoln(err)
			return
		}
		fileSize := fileHeader.Size

		srv := service.GetUserSrv()
		resp, err := srv.UserAvatarUpload(ctx.Request.Context(), file, fileSize, &req)
		if err != nil {
			//log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
	}
}

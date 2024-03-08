package api

import (
	"ChargPiles/pkg/e"
	"ChargPiles/pkg/utils/ctl"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, err error) *ctl.TrackedErrorResponse {
	//if ve, ok := err.(validator.ValidationErrors); ok {
	//	for _, fieldError := range ve {
	//		field := config.T(fmt.Sprintf("Field.%s", fieldError.Field))
	//		tag := config.T(fmt.Sprintf("Tag.Valid.%s", fieldError.Tag))
	//		return ctl.RespError(ctx, err, fmt.Sprintf("%s%s", field, tag))
	//	}
	//}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(ctx, err, "JSON类型不匹配")
	}

	return ctl.RespError(ctx, err, err.Error(), e.ERROR)
}

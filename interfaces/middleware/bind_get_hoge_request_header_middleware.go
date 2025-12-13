package middleware

import (
	"go_rest_api_template_with_gin/constant"
	v1 "go_rest_api_template_with_gin/interfaces/handler/v1"
	"go_rest_api_template_with_gin/interfaces/handler/v1/request"

	"github.com/gin-gonic/gin"
)

func BindGetHogeRequestHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := request.GetHogeRequestHeader{}
		if err := ctx.ShouldBindHeader(&header); err != nil {
			ctx.JSON(400,
				v1.CreateErrorResponse(
					header.XRequestID,
					err,
					constant.GP4100,
					"",
				))
			ctx.Abort()
		}
		ctx.Set("hogeID", header.HogeID)
	}
}

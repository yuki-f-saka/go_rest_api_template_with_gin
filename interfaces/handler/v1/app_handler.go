package v1

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"go_rest_api_template_with_gin/constant"
	"go_rest_api_template_with_gin/usecase"
)

type AppHandler struct {
	userUsecase usecase.UserUsecase
}

// NewAppHandler はアプリケーションハンドラーの新しいインスタンスを返します。
func NewAppHandler(userUsecase usecase.UserUsecase) *AppHandler {
	return &AppHandler{
		userUsecase: userUsecase,
	}
}

func CreateErrorResponse(requestId string, err error, errCode constant.ErrorCode, refUrl string) gin.H {
	return gin.H{
		"request_id": requestId,
		"timestamp":  time.Now().Format(constant.ZapDateTimelayout),
		"error": gin.H{
			"message": fmt.Sprintf("%v", err),
			"code":    errCode,
			"ref_urf": refUrl,
		},
	}
}

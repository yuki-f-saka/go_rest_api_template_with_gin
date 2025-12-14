package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"go_rest_api_template_with_gin/constant"
	"go_rest_api_template_with_gin/domain"
	"go_rest_api_template_with_gin/interfaces/handler/v1/request"
)

// PostUser はユーザー作成エンドポイント
func (h *AppHandler) PostUser(c *gin.Context) {
	var req request.PostUserRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CreateErrorResponse(
			c.GetString("request_id"),
			err,
			constant.GP4100,
			c.Request.RequestURI,
		))
		return
	}

	// Domain モデルにマップ
	user := &domain.User{
		ID:    uuid.New().String(),
		Name:  req.Name,
		Email: req.Email,
	}

	// ユースケースで作成
	if err := h.userUsecase.CreateUser(c.Request.Context(), user); err != nil {
		c.JSON(http.StatusInternalServerError, CreateErrorResponse(
			c.GetString("request_id"),
			err,
			constant.GP5000,
			c.Request.RequestURI,
		))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"request_id": c.GetString("request_id"),
		"data": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}

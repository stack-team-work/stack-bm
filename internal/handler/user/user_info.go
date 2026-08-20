package user

import (
	"net/http"

	userSvc "stack-bm/internal/service/user"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserInfoHandler struct {
	service *userSvc.UserInfoService
}

func NewUserInfoHandler() *UserInfoHandler {
	return &UserInfoHandler{service: userSvc.NewUserInfoService()}
}

func (h *UserInfoHandler) GetList(c *gin.Context) {
	var req struct {
		Page    int    `json:"page"`
		Size    int    `json:"size"`
		AppID   int    `json:"app_id"`
		UserID  string `json:"user_id"`
		StartAt int64  `json:"start_at"`
		EndAt   int64  `json:"end_at"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	list, total, err := h.service.FindPage(req.Page, req.Size, req.AppID, req.UserID, req.StartAt, req.EndAt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, req.Page, req.Size)
}
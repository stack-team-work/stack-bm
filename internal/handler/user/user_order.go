package user

import (
	"net/http"

	userSvc "stack-bm/internal/service/user"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserOrderHandler struct {
	service *userSvc.UserOrderService
}

func NewUserOrderHandler() *UserOrderHandler {
	return &UserOrderHandler{service: userSvc.NewUserOrderService()}
}

func (h *UserOrderHandler) GetList(c *gin.Context) {
	var req struct {
		Page      int    `json:"page"`
		Size      int    `json:"size"`
		AppID     int    `json:"app_id"`
		UserID    string `json:"user_id"`
		PayStatus int    `json:"pay_status"`
		Status    int    `json:"status"`
		StartAt   int64  `json:"start_at"`
		EndAt     int64  `json:"end_at"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	list, total, err := h.service.FindPage(req.Page, req.Size, req.AppID, req.UserID, req.PayStatus, req.Status, req.StartAt, req.EndAt)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, req.Page, req.Size)
}
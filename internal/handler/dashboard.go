package handler

import (
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type DashboardStats struct {
	RechargeUsers  int64 `json:"recharge_users"`
	RegisterUsers  int64 `json:"register_users"`
	LoginUsers     int64 `json:"login_users"`
	ActivateUsers  int64 `json:"activate_users"`
}

type DashboardHandler struct{}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

func (h *DashboardHandler) Stats(c *gin.Context) {
	stats := DashboardStats{
		RechargeUsers: 0,
		RegisterUsers: 0,
		LoginUsers:    0,
		ActivateUsers: 0,
	}

	response.Success(c, stats)
}

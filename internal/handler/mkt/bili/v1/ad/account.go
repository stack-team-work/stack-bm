package ad

import (
	"net/http"

	biliAdSvc "stack-bm/internal/service/mkt/bili/v1/ad"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// AccountHandler B站账户数据
type AccountHandler struct {
	service *biliAdSvc.AccountService
}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{service: biliAdSvc.NewAccountService()}
}

// Register 注册账户路由（挂载于 /bili/v1 组）
func (h *AccountHandler) Register(rg gin.IRouter) {
	rg.POST("/account/list", h.List)
}

// List 账户数据列表（支持指标列选择与筛选）
func (h *AccountHandler) List(c *gin.Context) {
	var req ListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	req.Normalize()
	list, total, err := h.service.List(req.Page, req.Size, req.Columns, req.Filters())
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.PageSuccess(c, list, total, req.Page, req.Size)
}

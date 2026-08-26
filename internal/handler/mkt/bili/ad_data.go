package bili

import (
	"net/http"

	biliSync "stack-bm/internal/service/mkt/bili/sync"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// AdDataHandler 广告数据
type AdDataHandler struct {
	service *biliSync.AdDataService
}

func NewAdDataHandler() *AdDataHandler {
	return &AdDataHandler{service: biliSync.NewAdDataService()}
}

// List 广告数据列表（按层级，支持指标列选择与筛选）
func (h *AdDataHandler) List(c *gin.Context) {
	level := c.Param("level")
	var req struct {
		Page      int                    `json:"page"`
		Size      int                    `json:"size"`
		Columns   []string               `json:"columns"`
		Keyword   string                 `json:"keyword"`
		Status    int                    `json:"status"`
		AccountID string                 `json:"account_id"`
		StartDate string                 `json:"start_date"`
		EndDate   string                 `json:"end_date"`
		Extra     map[string]interface{} `json:"extra"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 10
	}
	filters := map[string]interface{}{}
	if req.Keyword != "" {
		filters["keyword"] = req.Keyword
	}
	if req.Status != 0 {
		filters["status"] = req.Status
	}
	if req.AccountID != "" {
		filters["account_id"] = req.AccountID
	}
	if req.StartDate != "" {
		filters["start_date"] = req.StartDate
	}
	if req.EndDate != "" {
		filters["end_date"] = req.EndDate
	}
	for k, v := range req.Extra {
		filters[k] = v
	}
	list, total, err := h.service.List(level, req.Page, req.Size, req.Columns, filters)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.PageSuccess(c, list, total, req.Page, req.Size)
}
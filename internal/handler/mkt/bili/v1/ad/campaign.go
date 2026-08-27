package ad

import (
	"net/http"

	biliAdSvc "stack-bm/internal/service/mkt/bili/v1/ad"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// CampaignHandler B站计划（第一层级）
type CampaignHandler struct {
	service *biliAdSvc.CampaignService
}

func NewCampaignHandler() *CampaignHandler {
	return &CampaignHandler{service: biliAdSvc.NewCampaignService()}
}

// Register 注册计划路由（挂载于 /bili/v1 组）
func (h *CampaignHandler) Register(rg gin.IRouter) {
	rg.POST("/campaign/list", h.List)
	rg.POST("/campaign/open", h.Open)
	rg.POST("/campaign/pause", h.Pause)
	rg.POST("/campaign/delete", h.Delete)
	rg.POST("/campaign/set-budget", h.SetBudget)
	rg.POST("/campaign/set-bid", h.SetBid)
	rg.POST("/campaign/set-deep-bid", h.SetDeepBid)
	rg.POST("/campaign/batch-status", h.BatchStatus)
	rg.POST("/campaign/batch-delete", h.BatchDelete)
}

// List 计划列表（支持指标列选择与筛选）
func (h *CampaignHandler) List(c *gin.Context) {
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

func (h *CampaignHandler) bind(c *gin.Context) (*ActionRequest, bool) {
	var req ActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return nil, false
	}
	return &req, true
}

func (h *CampaignHandler) Open(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Open(req.ID)
	respond(c, result, err)
}

func (h *CampaignHandler) Pause(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Pause(req.ID)
	respond(c, result, err)
}

func (h *CampaignHandler) Delete(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Delete(req.ID)
	respond(c, result, err)
}

func (h *CampaignHandler) SetBudget(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetBudget(req.ID, req.Budget)
	respond(c, result, err)
}

func (h *CampaignHandler) SetBid(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetBid(req.ID, req.Bid)
	respond(c, result, err)
}

func (h *CampaignHandler) SetDeepBid(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetDeepBid(req.ID, req.DeepBid)
	respond(c, result, err)
}

func (h *CampaignHandler) BatchStatus(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.BatchUpdateStatus(req.IDs, req.Status)
	respond(c, result, err)
}

func (h *CampaignHandler) BatchDelete(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.BatchDelete(req.IDs)
	respond(c, result, err)
}

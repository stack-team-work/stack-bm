package ad

import (
	"net/http"

	ksAdSvc "stack-bm/internal/service/mkt/ks/v1/ad"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// CreativeHandler B站创意（第三层级）
type CreativeHandler struct {
	service *ksAdSvc.CreativeService
}

func NewCreativeHandler() *CreativeHandler {
	return &CreativeHandler{service: ksAdSvc.NewCreativeService()}
}

// Register 注册创意路由（挂载于 /bili/v1 组）
func (h *CreativeHandler) Register(rg gin.IRouter) {
	rg.POST("/creative/list", h.List)
	rg.POST("/creative/open", h.Open)
	rg.POST("/creative/pause", h.Pause)
	rg.POST("/creative/preview", h.Preview)
	rg.POST("/creative/batch-status", h.BatchStatus)
}

// List 创意列表（支持指标列选择与筛选）
func (h *CreativeHandler) List(c *gin.Context) {
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

func (h *CreativeHandler) bind(c *gin.Context) (*ActionRequest, bool) {
	var req ActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return nil, false
	}
	return &req, true
}

func (h *CreativeHandler) Open(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Open(req.ID)
	respond(c, result, err)
}

func (h *CreativeHandler) Pause(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Pause(req.ID)
	respond(c, result, err)
}

func (h *CreativeHandler) Preview(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Preview(req.ID)
	respond(c, result, err)
}

func (h *CreativeHandler) BatchStatus(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.BatchUpdateStatus(req.IDs, req.Status)
	respond(c, result, err)
}

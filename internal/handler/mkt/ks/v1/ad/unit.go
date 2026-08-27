package ad

import (
	"net/http"

	ksAdSvc "stack-bm/internal/service/mkt/ks/v1/ad"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// UnitHandler B站单元（第二层级）
type UnitHandler struct {
	service *ksAdSvc.UnitService
}

func NewUnitHandler() *UnitHandler {
	return &UnitHandler{service: ksAdSvc.NewUnitService()}
}

// Register 注册单元路由（挂载于 /bili/v1 组）
func (h *UnitHandler) Register(rg gin.IRouter) {
	rg.POST("/unit/list", h.List)
	rg.POST("/unit/open", h.Open)
	rg.POST("/unit/pause", h.Pause)
	rg.POST("/unit/delete", h.Delete)
	rg.POST("/unit/set-budget", h.SetBudget)
	rg.POST("/unit/set-bid", h.SetBid)
	rg.POST("/unit/set-deep-bid", h.SetDeepBid)
	rg.POST("/unit/set-begin-date", h.SetBeginDate)
	rg.POST("/unit/collect", h.Collect)
	rg.POST("/unit/cancel-collect", h.CancelCollect)
	rg.POST("/unit/set-raise", h.SetRaise)
	rg.POST("/unit/stop-raise", h.StopRaise)
	rg.POST("/unit/raise-info", h.RaiseInfo)
	rg.POST("/unit/batch-status", h.BatchStatus)
	rg.POST("/unit/batch-delete", h.BatchDelete)
	rg.POST("/unit/batch-set-raise", h.BatchSetRaise)
	rg.POST("/unit/batch-stop-raise", h.BatchStopRaise)
}

// List 单元列表（支持指标列选择与筛选）
func (h *UnitHandler) List(c *gin.Context) {
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

func (h *UnitHandler) bind(c *gin.Context) (*ActionRequest, bool) {
	var req ActionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return nil, false
	}
	return &req, true
}

func respond(c *gin.Context, result interface{}, err error) {
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, result)
}

func (h *UnitHandler) Open(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Open(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) Pause(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Pause(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) Delete(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Delete(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) SetBudget(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetBudget(req.ID, req.Budget)
	respond(c, result, err)
}

func (h *UnitHandler) SetBid(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetBid(req.ID, req.Bid)
	respond(c, result, err)
}

func (h *UnitHandler) SetDeepBid(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetDeepBid(req.ID, req.DeepBid)
	respond(c, result, err)
}

func (h *UnitHandler) SetBeginDate(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetBeginDate(req.ID, req.BeginDate)
	respond(c, result, err)
}

func (h *UnitHandler) Collect(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.Collect(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) CancelCollect(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.CancelCollect(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) SetRaise(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.SetRaise(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) StopRaise(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.StopRaise(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) RaiseInfo(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.RaiseInfo(req.ID)
	respond(c, result, err)
}

func (h *UnitHandler) BatchStatus(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.BatchUpdateStatus(req.IDs, req.Status)
	respond(c, result, err)
}

func (h *UnitHandler) BatchDelete(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.BatchDelete(req.IDs)
	respond(c, result, err)
}

func (h *UnitHandler) BatchSetRaise(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.BatchSetRaise(req.IDs)
	respond(c, result, err)
}

func (h *UnitHandler) BatchStopRaise(c *gin.Context) {
	req, ok := h.bind(c)
	if !ok {
		return
	}
	result, err := h.service.BatchStopRaise(req.IDs)
	respond(c, result, err)
}

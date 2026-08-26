package ks

import (
	"errors"
	"net/http"

	ksTool "stack-bm/internal/service/mkt/ks/tool"
	"stack-bm/internal/service/mkt/oauth"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// ToolHandler 批量操作（按层级+action）
type ToolHandler struct {
	campaign *ksTool.CampaignTool
	unit     *ksTool.UnitTool
	creative *ksTool.CreativeTool
}

func NewToolHandler() *ToolHandler {
	auth := oauth.NewManagerAuth()
	return &ToolHandler{
		campaign: ksTool.NewCampaignTool(auth),
		unit:     ksTool.NewUnitTool(auth),
		creative: ksTool.NewCreativeTool(auth),
	}
}

// Action 批量操作入口
func (h *ToolHandler) Action(c *gin.Context) {
	level := c.Param("level")
	action := c.Param("action")
	var req ksTool.ToolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	var result interface{}
	var err error
	switch level {
	case "campaign":
		result, err = h.dispatchCampaign(action, &req)
	case "unit":
		result, err = h.dispatchUnit(action, &req)
	case "creative":
		result, err = h.dispatchCreative(action, &req)
	default:
		response.Error(c, http.StatusBadRequest, "未知层级: "+level)
		return
	}
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, result)
}

func (h *ToolHandler) dispatchCampaign(action string, req *ksTool.ToolRequest) (interface{}, error) {
	switch action {
	case "open":
		return h.campaign.Open(req.ID)
	case "pause":
		return h.campaign.Pause(req.ID)
	case "delete":
		return h.campaign.Delete(req.ID)
	case "set-budget":
		return h.campaign.SetBudget(req.ID, req.Budget)
	case "set-bid":
		return h.campaign.SetBid(req.ID, req.Bid)
	case "set-deep-bid":
		return h.campaign.SetDeepBid(req.ID, req.DeepBid)
	case "batch-status":
		return h.campaign.BatchUpdateStatus(req.IDs, req.Status)
	case "batch-delete":
		return h.campaign.BatchDelete(req.IDs)
	default:
		return nil, errors.New("未知操作: " + action)
	}
}

func (h *ToolHandler) dispatchUnit(action string, req *ksTool.ToolRequest) (interface{}, error) {
	switch action {
	case "open":
		return h.unit.Open(req.ID)
	case "pause":
		return h.unit.Pause(req.ID)
	case "delete":
		return h.unit.Delete(req.ID)
	case "set-budget":
		return h.unit.SetBudget(req.ID, req.Budget)
	case "set-bid":
		return h.unit.SetBid(req.ID, req.Bid)
	case "set-deep-bid":
		return h.unit.SetDeepBid(req.ID, req.DeepBid)
	case "set-begin-date":
		return h.unit.SetBeginDate(req.ID, req.BeginDate)
	case "batch-status":
		return h.unit.BatchUpdateStatus(req.IDs, req.Status)
	case "batch-delete":
		return h.unit.BatchDelete(req.IDs)
	case "collect":
		return h.unit.Collect(req.ID)
	case "cancel-collect":
		return h.unit.CancelCollect(req.ID)
	case "set-raise":
		return h.unit.SetRaise(req.ID)
	case "stop-raise":
		return h.unit.StopRaise(req.ID)
	case "batch-set-raise":
		return h.unit.BatchSetRaise(req.IDs)
	case "batch-stop-raise":
		return h.unit.BatchStopRaise(req.IDs)
	case "raise-info":
		return h.unit.RaiseInfo(req.ID)
	default:
		return nil, errors.New("未知操作: " + action)
	}
}

func (h *ToolHandler) dispatchCreative(action string, req *ksTool.ToolRequest) (interface{}, error) {
	switch action {
	case "open":
		return h.creative.Open(req.ID)
	case "pause":
		return h.creative.Pause(req.ID)
	case "preview":
		return h.creative.Preview(req.ID)
	case "batch-status":
		return h.creative.BatchUpdateStatus(req.IDs, req.Status)
	default:
		return nil, errors.New("未知操作: " + action)
	}
}
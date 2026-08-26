package tt

import (
	"errors"
	"net/http"

	"stack-bm/internal/service/mkt/oauth"
	ttTool "stack-bm/internal/service/mkt/tt/tool"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// ToolHandler 批量操作（按层级+action）
type ToolHandler struct {
	campaign *ttTool.CampaignTool
}

func NewToolHandler() *ToolHandler {
	auth := oauth.NewManagerAuth()
	return &ToolHandler{campaign: ttTool.NewCampaignTool(auth)}
}

// Action 批量操作入口
func (h *ToolHandler) Action(c *gin.Context) {
	level := c.Param("level")
	action := c.Param("action")
	var req ttTool.ToolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	var result interface{}
	var err error
	switch level {
	case "campaign":
		result, err = h.dispatchCampaign(action, &req)
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

func (h *ToolHandler) dispatchCampaign(action string, req *ttTool.ToolRequest) (interface{}, error) {
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
package handler

import (
	"net/http"

	gameSvc "stack-bm/internal/service/sdk/game"
	paySvc "stack-bm/internal/service/sdk/pay"
	mediaSvc "stack-bm/internal/service/mkt/media"
	feishuSvc "stack-bm/internal/service/bm/feishu"
	sysadminSvc "stack-bm/internal/service/bm/sys"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type OptionsHandler struct {
	providers map[string]func() ([]dict.Option, error)
}

func NewOptionsHandler() *OptionsHandler {
	adminSvc := sysadminSvc.NewSysAdminService()
	gameService := gameSvc.NewGameService()
	payService := paySvc.NewPayPlatformService()
	mediaService := mediaSvc.NewMediaService()

	return &OptionsHandler{
		providers: map[string]func() ([]dict.Option, error){
			"admin":              adminSvc.FindOptions,
			"feishu_app":         feishuSvc.NewFeishuAppService().FindOptions,
			"game":               gameService.FindOptions,
			"game_app":          gameSvc.NewGameAppService().FindOptions,
			"game_app_template":  gameSvc.NewGameAppTemplateService().FindOptions,
			"game_cp":            gameSvc.NewGameCpService().FindOptions,
			"game_gift":          gameSvc.NewGameGiftService().FindOptions,
			"game_voucher":       gameSvc.NewGameVoucherService().FindOptions,
			"pay_platform":       payService.FindOptions,
			"media":              mediaService.FindOptions,
			"media_subject":      mediaSvc.NewMediaSubjectService().FindOptions,
			"media_application":  mediaSvc.NewMediaApplicationService().FindOptions,
			"media_sub":          mediaSvc.NewMediaSubService().FindOptions,
			"media_agent":        mediaSvc.NewMediaAgentService().FindOptions,
			"media_manager":      mediaSvc.NewMediaManagerService().FindOptions,
			"media_account":      mediaSvc.NewMediaAccountService().FindOptions,
		},
	}
}

func (h *OptionsHandler) Get(c *gin.Context) {
	var req struct{ Type string `json:"type"` }
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	provider, ok := h.providers[req.Type]
	if !ok {
		response.Error(c, http.StatusNotFound, "未知的option类型: "+req.Type)
		return
	}
	list, err := provider()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

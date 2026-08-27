package router

import (
	"stack-bm/internal/handler"
	"stack-bm/internal/handler/bm/feishu"
	"stack-bm/internal/handler/bm/sys"
	biliadh "stack-bm/internal/handler/mkt/bili/v1/ad"
	bilitplh "stack-bm/internal/handler/mkt/bili/v1/template"
	ksadh "stack-bm/internal/handler/mkt/ks/v1/ad"
	kstplh "stack-bm/internal/handler/mkt/ks/v1/template"
	"stack-bm/internal/handler/mkt/media"
	tcadh "stack-bm/internal/handler/mkt/tc/v1/ad"
	ttadh "stack-bm/internal/handler/mkt/tt/v1/ad"
	tttplh "stack-bm/internal/handler/mkt/tt/v1/template"
	"stack-bm/internal/handler/sdk/game"
	"stack-bm/internal/handler/sdk/pay"
	sdkSys "stack-bm/internal/handler/sdk/sys"
	"stack-bm/internal/handler/user"
	"stack-bm/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggingMiddleware())

	authHandler := handler.NewAuthHandler()
	dashboardHandler := handler.NewDashboardHandler()
	dictHandler := handler.NewDictHandler()
	optionsHandler := handler.NewOptionsHandler()
	sysAdminHandler := sys.NewSysAdminHandler()
	sysAdminGroupHandler := sys.NewSysAdminGroupHandler()
	sysLogHandler := sys.NewSysLogHandler()
	sysMenuHandler := sys.NewSysMenuHandler()
	sysColumnHandler := sys.NewSysColumnHandler()
	sysTagHandler := sys.NewSysTagHandler()

	feishuAppHandler := feishu.NewFeishuAppHandler()
	feishuChatHandler := feishu.NewFeishuChatHandler()
	feishuUserHandler := feishu.NewFeishuUserHandler()
	gameHandler := game.NewGameHandler()
	gameAppHandler := game.NewGameAppHandler()
	gameAppTemplateHandler := game.NewGameAppTemplateHandler()
	gameCpHandler := game.NewGameCpHandler()
	gameTagHandler := game.NewGameTagHandler()
	gameVariableHandler := game.NewGameVariableHandler()
	gamePlatformHandler := game.NewGamePlatformHandler()
	gameGiftHandler := game.NewGameGiftHandler()
	gameGiftCodeHandler := game.NewGameGiftCodeHandler()
	gameGiftUserCodeHandler := game.NewGameGiftUserCodeHandler()
	gameVoucherHandler := game.NewGameVoucherHandler()
	gameVoucherUseHandler := game.NewGameVoucherUseHandler()
	sdkSysLogHandler := sdkSys.NewSysLogHandler()
	mediaHandler := media.NewMediaHandler()
	mediaSubHandler := media.NewMediaSubHandler()
	mediaAgentHandler := media.NewMediaAgentHandler()
	mediaApplicationHandler := media.NewMediaApplicationHandler()
	mediaManagerHandler := media.NewMediaManagerHandler()
	mediaSubjectHandler := media.NewMediaSubjectHandler()
	mediaAccountHandler := media.NewMediaAccountHandler()
	oauthCallbackHandler := media.NewOAuthCallbackHandler()

	payPlatformHandler := pay.NewPayPlatformHandler()
	payMerchantHandler := pay.NewPayMerchantHandler()

	userInfoHandler := user.NewUserInfoHandler()
	userOrderHandler := user.NewUserOrderHandler()
	userLoginHandler := user.NewUserLoginHandler()
	userActiveHandler := user.NewUserActiveHandler()

	r.POST("/api/login", authHandler.Login)
	r.POST("/api/captcha", authHandler.Captcha)
	r.POST("/api/dict", dictHandler.GetAll)
	r.POST("/api/dict/:key", dictHandler.GetByKey)

	r.GET("/oauth/callback/:channel", oauthCallbackHandler.Callback)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/user/info", authHandler.GetUserInfo)
		api.POST("/options", optionsHandler.Get)

		api.POST("/dashboard/stats", dashboardHandler.Stats)

		api.POST("/admin/create", sysAdminHandler.Create)
		api.POST("/admin/list", sysAdminHandler.GetList)
		api.POST("/admin/detail/:id", sysAdminHandler.GetByID)
		api.POST("/admin/update/:id", sysAdminHandler.Update)
		api.POST("/admin/delete/:id", sysAdminHandler.Delete)

		api.POST("/admin-group/create", sysAdminGroupHandler.Create)
		api.POST("/admin-group/list", sysAdminGroupHandler.GetList)
		api.POST("/admin-group/all", sysAdminGroupHandler.GetAll)
		api.POST("/admin-group/detail/:id", sysAdminGroupHandler.GetByID)
		api.POST("/admin-group/update/:id", sysAdminGroupHandler.Update)
		api.POST("/admin-group/delete/:id", sysAdminGroupHandler.Delete)

		api.POST("/game/create", gameHandler.Create)
		api.POST("/game/list", gameHandler.GetList)
		api.POST("/game/all", gameHandler.GetAll)
		api.POST("/game/detail/:id", gameHandler.GetByID)
		api.POST("/game/update/:id", gameHandler.Update)
		api.POST("/game/delete/:id", gameHandler.Delete)

		api.POST("/game-app/create", gameAppHandler.Create)
		api.POST("/game-app/list", gameAppHandler.GetList)
		api.POST("/game-app/all", gameAppHandler.GetAll)
		api.POST("/game-app/detail/:id", gameAppHandler.GetByID)
		api.POST("/game-app/update/:id", gameAppHandler.Update)
		api.POST("/game-app/delete/:id", gameAppHandler.Delete)

		api.POST("/game-app-template/create", gameAppTemplateHandler.Create)
		api.POST("/game-app-template/list", gameAppTemplateHandler.GetList)
		api.POST("/game-app-template/all", gameAppTemplateHandler.GetAll)
		api.POST("/game-app-template/detail/:id", gameAppTemplateHandler.GetByID)
		api.POST("/game-app-template/update/:id", gameAppTemplateHandler.Update)
		api.POST("/game-app-template/delete/:id", gameAppTemplateHandler.Delete)

		api.POST("/game-cp/create", gameCpHandler.Create)
		api.POST("/game-cp/list", gameCpHandler.GetList)
		api.POST("/game-cp/all", gameCpHandler.GetAll)
		api.POST("/game-cp/detail/:id", gameCpHandler.GetByID)
		api.POST("/game-cp/update/:id", gameCpHandler.Update)
		api.POST("/game-cp/delete/:id", gameCpHandler.Delete)

		api.POST("/game-tag/create", gameTagHandler.Create)
		api.POST("/game-tag/list", gameTagHandler.GetList)
		api.POST("/game-tag/all", gameTagHandler.GetAll)
		api.POST("/game-tag/detail/:id", gameTagHandler.GetByID)
		api.POST("/game-tag/update/:id", gameTagHandler.Update)
		api.POST("/game-tag/delete/:id", gameTagHandler.Delete)

		api.POST("/game-variable/create", gameVariableHandler.Create)
		api.POST("/game-variable/list", gameVariableHandler.GetList)
		api.POST("/game-variable/detail/:id", gameVariableHandler.GetByID)
		api.POST("/game-variable/update/:id", gameVariableHandler.Update)
		api.POST("/game-variable/delete/:id", gameVariableHandler.Delete)

		api.POST("/game-platform/create", gamePlatformHandler.Create)
		api.POST("/game-platform/list", gamePlatformHandler.GetList)
		api.POST("/game-platform/all", gamePlatformHandler.GetAll)
		api.POST("/game-platform/detail/:id", gamePlatformHandler.GetByID)
		api.POST("/game-platform/update/:id", gamePlatformHandler.Update)
		api.POST("/game-platform/delete/:id", gamePlatformHandler.Delete)

		api.POST("/game-gift/create", gameGiftHandler.Create)
		api.POST("/game-gift/list", gameGiftHandler.GetList)
		api.POST("/game-gift/all", gameGiftHandler.GetAll)
		api.POST("/game-gift/detail/:id", gameGiftHandler.GetByID)
		api.POST("/game-gift/update/:id", gameGiftHandler.Update)
		api.POST("/game-gift/delete/:id", gameGiftHandler.Delete)

		api.POST("/game-gift-code/create", gameGiftCodeHandler.Create)
		api.POST("/game-gift-code/list", gameGiftCodeHandler.GetList)
		api.POST("/game-gift-code/detail/:id", gameGiftCodeHandler.GetByID)
		api.POST("/game-gift-code/update/:id", gameGiftCodeHandler.Update)
		api.POST("/game-gift-code/delete/:id", gameGiftCodeHandler.Delete)

		api.POST("/game-gift-user-code/list", gameGiftUserCodeHandler.GetList)

		api.POST("/game-voucher/create", gameVoucherHandler.Create)
		api.POST("/game-voucher/list", gameVoucherHandler.GetList)
		api.POST("/game-voucher/all", gameVoucherHandler.GetAll)
		api.POST("/game-voucher/detail/:id", gameVoucherHandler.GetByID)
		api.POST("/game-voucher/update/:id", gameVoucherHandler.Update)
		api.POST("/game-voucher/delete/:id", gameVoucherHandler.Delete)

		api.POST("/game-voucher-use/list", gameVoucherUseHandler.GetList)

		api.POST("/logs/list", sysLogHandler.GetList)
		api.POST("/logs/clear", sysLogHandler.ClearAll)

		api.POST("/menu/create", sysMenuHandler.Create)
		api.POST("/menu/list", sysMenuHandler.GetList)
		api.POST("/menu/all", sysMenuHandler.GetAll)
		api.POST("/menu/detail/:id", sysMenuHandler.GetByID)
		api.POST("/menu/update/:id", sysMenuHandler.Update)
		api.POST("/menu/delete/:id", sysMenuHandler.Delete)

		api.POST("/sys-column/create", sysColumnHandler.Create)
		api.POST("/sys-column/list", sysColumnHandler.GetList)
		api.POST("/sys-column/all", sysColumnHandler.GetAll)
		api.POST("/sys-column/detail/:id", sysColumnHandler.GetByID)
		api.POST("/sys-column/update/:id", sysColumnHandler.Update)
		api.POST("/sys-column/delete/:id", sysColumnHandler.Delete)

		api.POST("/sys-tag/create", sysTagHandler.Create)
		api.POST("/sys-tag/list", sysTagHandler.GetList)
		api.POST("/sys-tag/all", sysTagHandler.GetAll)
		api.POST("/sys-tag/detail/:id", sysTagHandler.GetByID)
		api.POST("/sys-tag/update/:id", sysTagHandler.Update)

		api.POST("/feishu-app/create", feishuAppHandler.Create)
		api.POST("/feishu-app/list", feishuAppHandler.GetList)
		api.POST("/feishu-app/all", feishuAppHandler.GetAll)
		api.POST("/feishu-app/detail/:id", feishuAppHandler.GetByID)
		api.POST("/feishu-app/update/:id", feishuAppHandler.Update)
		api.POST("/feishu-app/status/:id", feishuAppHandler.UpdateStatus)

		api.POST("/feishu-chat/create", feishuChatHandler.Create)
		api.POST("/feishu-chat/list", feishuChatHandler.GetList)
		api.POST("/feishu-chat/all", feishuChatHandler.GetAll)
		api.POST("/feishu-chat/detail/:id", feishuChatHandler.GetByID)
		api.POST("/feishu-chat/update/:id", feishuChatHandler.Update)
		api.POST("/feishu-chat/status/:id", feishuChatHandler.UpdateStatus)

		api.POST("/feishu-user/create", feishuUserHandler.Create)
		api.POST("/feishu-user/list", feishuUserHandler.GetList)
		api.POST("/feishu-user/all", feishuUserHandler.GetAll)
		api.POST("/feishu-user/detail/:id", feishuUserHandler.GetByID)
		api.POST("/feishu-user/update/:id", feishuUserHandler.Update)
		api.POST("/feishu-user/status/:id", feishuUserHandler.UpdateStatus)

		api.POST("/sdk-logs/list", sdkSysLogHandler.GetList)

		api.POST("/media/create", mediaHandler.Create)
		api.POST("/media/list", mediaHandler.GetList)
		api.POST("/media/all", mediaHandler.GetAll)
		api.POST("/media/detail/:id", mediaHandler.GetByID)
		api.POST("/media/update/:id", mediaHandler.Update)
		api.POST("/media/delete/:id", mediaHandler.Delete)

		api.POST("/media-sub/create", mediaSubHandler.Create)
		api.POST("/media-sub/list", mediaSubHandler.GetList)
		api.POST("/media-sub/all", mediaSubHandler.GetAll)
		api.POST("/media-sub/detail/:id", mediaSubHandler.GetByID)
		api.POST("/media-sub/update/:id", mediaSubHandler.Update)
		api.POST("/media-sub/delete/:id", mediaSubHandler.Delete)

		api.POST("/media-agent/create", mediaAgentHandler.Create)
		api.POST("/media-agent/list", mediaAgentHandler.GetList)
		api.POST("/media-agent/all", mediaAgentHandler.GetAll)
		api.POST("/media-agent/detail/:id", mediaAgentHandler.GetByID)
		api.POST("/media-agent/update/:id", mediaAgentHandler.Update)
		api.POST("/media-agent/delete/:id", mediaAgentHandler.Delete)

		api.POST("/media-application/create", mediaApplicationHandler.Create)
		api.POST("/media-application/list", mediaApplicationHandler.GetList)
		api.POST("/media-application/all", mediaApplicationHandler.GetAll)
		api.POST("/media-application/detail/:id", mediaApplicationHandler.GetByID)
		api.POST("/media-application/update/:id", mediaApplicationHandler.Update)
		api.POST("/media-application/delete/:id", mediaApplicationHandler.Delete)

		api.POST("/media-manager/create", mediaManagerHandler.Create)
		api.POST("/media-manager/list", mediaManagerHandler.GetList)
		api.POST("/media-manager/all", mediaManagerHandler.GetAll)
		api.POST("/media-manager/detail/:id", mediaManagerHandler.GetByID)
		api.POST("/media-manager/update/:id", mediaManagerHandler.Update)
		api.POST("/media-manager/delete/:id", mediaManagerHandler.Delete)
		api.POST("/media-manager/oauth", mediaManagerHandler.Oauth)
		api.POST("/media-manager/sync-advertiser", mediaManagerHandler.SyncAdvertiser)

		api.POST("/media-subject/create", mediaSubjectHandler.Create)
		api.POST("/media-subject/list", mediaSubjectHandler.GetList)
		api.POST("/media-subject/all", mediaSubjectHandler.GetAll)
		api.POST("/media-subject/detail/:id", mediaSubjectHandler.GetByID)
		api.POST("/media-subject/update/:id", mediaSubjectHandler.Update)
		api.POST("/media-subject/delete/:id", mediaSubjectHandler.Delete)

		api.POST("/media-account/create", mediaAccountHandler.Create)
		api.POST("/media-account/list", mediaAccountHandler.GetList)
		api.POST("/media-account/all", mediaAccountHandler.GetAll)
		api.POST("/media-account/detail/:id", mediaAccountHandler.GetByID)
		api.POST("/media-account/update/:id", mediaAccountHandler.Update)
		api.POST("/media-account/delete/:id", mediaAccountHandler.Delete)

		// 四渠道广告数据 v1（版本为本项目自定义：头条 v1 对应巨量V3，B站 v1 对应当前开放API）
		biliV1 := api.Group("/bili/v1")
		biliadh.RegisterRoutes(biliV1)
		bilitplh.RegisterRoutes(biliV1)

		ksV1 := api.Group("/ks/v1")
		ksadh.RegisterRoutes(ksV1)
		kstplh.RegisterRoutes(ksV1)

		ttV1 := api.Group("/tt/v1")
		ttadh.RegisterRoutes(ttV1)
		tttplh.RegisterRoutes(ttV1)

		tcV1 := api.Group("/tc/v1")
		tcadh.RegisterRoutes(tcV1)

		api.POST("/pay-platform/create", payPlatformHandler.Create)
		api.POST("/pay-platform/list", payPlatformHandler.GetList)
		api.POST("/pay-platform/all", payPlatformHandler.GetAll)
		api.POST("/pay-platform/detail/:id", payPlatformHandler.GetByID)
		api.POST("/pay-platform/update/:id", payPlatformHandler.Update)
		api.POST("/pay-platform/delete/:id", payPlatformHandler.Delete)

		api.POST("/pay-merchant/create", payMerchantHandler.Create)
		api.POST("/pay-merchant/list", payMerchantHandler.GetList)
		api.POST("/pay-merchant/all", payMerchantHandler.GetAll)
		api.POST("/pay-merchant/detail/:id", payMerchantHandler.GetByID)
		api.POST("/pay-merchant/update/:id", payMerchantHandler.Update)
		api.POST("/pay-merchant/delete/:id", payMerchantHandler.Delete)

		api.POST("/user/info/list", userInfoHandler.GetList)
		api.POST("/user/orders/list", userOrderHandler.GetList)
		api.POST("/user/logins/list", userLoginHandler.GetList)
		api.POST("/user/actives/list", userActiveHandler.GetList)
	}

	return r
}

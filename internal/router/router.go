package router

import (
	"stack-bm/internal/handler"
	"stack-bm/internal/handler/bm/sys"
	"stack-bm/internal/handler/mkt/media"
	"stack-bm/internal/handler/sdk/game"
	sdkSys "stack-bm/internal/handler/sdk/sys"
	"stack-bm/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggingMiddleware())

	authHandler := handler.NewAuthHandler()
	dashboardHandler := handler.NewDashboardHandler()
	sysAdminHandler := sys.NewSysAdminHandler()
	sysAdminGroupHandler := sys.NewSysAdminGroupHandler()
	sysLogHandler := sys.NewSysLogHandler()
	sysMenuHandler := sys.NewSysMenuHandler()
	gameHandler := game.NewGameHandler()
	gameAppHandler := game.NewGameAppHandler()
	gameCpHandler := game.NewGameCpHandler()
	gameTagHandler := game.NewGameTagHandler()
	gameVariableHandler := game.NewGameVariableHandler()
	gamePlatformHandler := game.NewGamePlatformHandler()
	sdkSysLogHandler := sdkSys.NewSysLogHandler()
	mediaHandler := media.NewMediaHandler()
	mediaSubHandler := media.NewMediaSubHandler()
	mediaAgentHandler := media.NewMediaAgentHandler()
	mediaApplicationHandler := media.NewMediaApplicationHandler()
	mediaManagerHandler := media.NewMediaManagerHandler()
	mediaSubjectHandler := media.NewMediaSubjectHandler()

	r.POST("/api/login", authHandler.Login)
	r.POST("/api/captcha", authHandler.Captcha)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/user/info", authHandler.GetUserInfo)

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

		api.POST("/logs/list", sysLogHandler.GetList)
		api.POST("/logs/clear", sysLogHandler.ClearAll)

		api.POST("/menu/create", sysMenuHandler.Create)
		api.POST("/menu/list", sysMenuHandler.GetList)
		api.POST("/menu/all", sysMenuHandler.GetAll)
		api.POST("/menu/detail/:id", sysMenuHandler.GetByID)
		api.POST("/menu/update/:id", sysMenuHandler.Update)
		api.POST("/menu/delete/:id", sysMenuHandler.Delete)

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

		api.POST("/media-subject/create", mediaSubjectHandler.Create)
		api.POST("/media-subject/list", mediaSubjectHandler.GetList)
		api.POST("/media-subject/all", mediaSubjectHandler.GetAll)
		api.POST("/media-subject/detail/:id", mediaSubjectHandler.GetByID)
		api.POST("/media-subject/update/:id", mediaSubjectHandler.Update)
		api.POST("/media-subject/delete/:id", mediaSubjectHandler.Delete)
	}

	return r
}

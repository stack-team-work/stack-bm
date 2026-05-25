package router

import (
	"stack-bm/internal/handler"
	handlerGame "stack-bm/internal/handler/game"
	handlerSys "stack-bm/internal/handler/sys"
	"stack-bm/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	authHandler := handler.NewAuthHandler()
	sysAdminHandler := handlerSys.NewSysAdminHandler()
	sysAdminGroupHandler := handlerSys.NewSysAdminGroupHandler()
	gameHandler := handlerGame.NewGameHandler()
	gameAppHandler := handlerGame.NewGameAppHandler()
	gameCpHandler := handlerGame.NewGameCpHandler()

	r.POST("/api/login", authHandler.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/user/info", authHandler.GetUserInfo)

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
		api.POST("/game-app/detail/:id", gameAppHandler.GetByID)
		api.POST("/game-app/update/:id", gameAppHandler.Update)
		api.POST("/game-app/delete/:id", gameAppHandler.Delete)

		api.POST("/game-cp/create", gameCpHandler.Create)
		api.POST("/game-cp/list", gameCpHandler.GetList)
		api.POST("/game-cp/all", gameCpHandler.GetAll)
		api.POST("/game-cp/detail/:id", gameCpHandler.GetByID)
		api.POST("/game-cp/update/:id", gameCpHandler.Update)
		api.POST("/game-cp/delete/:id", gameCpHandler.Delete)
	}

	return r
}

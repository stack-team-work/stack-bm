package game

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sdk/game"
	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type GameAppHandler struct {
	service *gameSvc.GameAppService
}

func NewGameAppHandler() *GameAppHandler { return &GameAppHandler{service: gameSvc.NewGameAppService()} }

func (h *GameAppHandler) Create(c *gin.Context) {
	var app game.GameApp
	if err := c.ShouldBindJSON(&app); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if app.Name == "" || app.Pid == 0 {
		response.Error(c, http.StatusBadRequest, "应用名称和游戏ID不能为空")
		return
	}
	if err := h.service.Create(&app); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, app)
}

func (h *GameAppHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	gameID, _ := strconv.Atoi(c.DefaultPostForm("game_id", "0"))
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	apps, total, err := h.service.FindPage(page, size, keyword, gameID, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, apps, total, page, size)
}

func (h *GameAppHandler) GetAll(c *gin.Context) {
	apps, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, apps)
}

func (h *GameAppHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	app, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "应用不存在")
		return
	}
	response.Success(c, app)
}

func (h *GameAppHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var app game.GameApp
	if err := c.ShouldBindJSON(&app); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &app); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *GameAppHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

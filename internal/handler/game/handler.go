package game

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/game"
	gameSvc "stack-bm/internal/service/game"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	service *gameSvc.GameService
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		service: gameSvc.NewGameService(),
	}
}

func (h *GameHandler) Create(c *gin.Context) {
	var g game.Game
	if err := c.ShouldBindJSON(&g); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if g.Name == "" || g.Mark == "" {
		response.Error(c, http.StatusBadRequest, "游戏名称和标识不能为空")
		return
	}

	if err := h.service.Create(&g); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, g)
}

func (h *GameHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))

	games, total, err := h.service.FindPage(page, size, keyword, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.PageSuccess(c, games, total, page, size)
}

func (h *GameHandler) GetAll(c *gin.Context) {
	games, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, games)
}

func (h *GameHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	g, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "游戏不存在")
		return
	}

	response.Success(c, g)
}

func (h *GameHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	var g game.Game
	if err := c.ShouldBindJSON(&g); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if err := h.service.Update(uint(id), &g); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *GameHandler) Delete(c *gin.Context) {
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

type GameAppHandler struct {
	service *gameSvc.GameAppService
}

func NewGameAppHandler() *GameAppHandler {
	return &GameAppHandler{
		service: gameSvc.NewGameAppService(),
	}
}

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

type GameCpHandler struct {
	service *gameSvc.GameCpService
}

func NewGameCpHandler() *GameCpHandler {
	return &GameCpHandler{
		service: gameSvc.NewGameCpService(),
	}
}

func (h *GameCpHandler) Create(c *gin.Context) {
	var cp game.GameCp
	if err := c.ShouldBindJSON(&cp); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if cp.Name == "" {
		response.Error(c, http.StatusBadRequest, "CP名称不能为空")
		return
	}

	if err := h.service.Create(&cp); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, cp)
}

func (h *GameCpHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))

	cps, total, err := h.service.FindPage(page, size, keyword, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.PageSuccess(c, cps, total, page, size)
}

func (h *GameCpHandler) GetAll(c *gin.Context) {
	cps, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, cps)
}

func (h *GameCpHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	cp, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "CP不存在")
		return
	}

	response.Success(c, cp)
}

func (h *GameCpHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	var cp game.GameCp
	if err := c.ShouldBindJSON(&cp); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if err := h.service.Update(uint(id), &cp); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *GameCpHandler) Delete(c *gin.Context) {
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

type GameTagHandler struct {
	service *gameSvc.GameTagService
}

func NewGameTagHandler() *GameTagHandler {
	return &GameTagHandler{
		service: gameSvc.NewGameTagService(),
	}
}

func (h *GameTagHandler) Create(c *gin.Context) {
	var tag game.GameTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if tag.Name == "" || tag.Mark == "" {
		response.Error(c, http.StatusBadRequest, "标签名称和标识不能为空")
		return
	}

	if err := h.service.Create(&tag); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, tag)
}

func (h *GameTagHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	tagType, _ := strconv.Atoi(c.DefaultPostForm("type", "0"))
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))

	tags, total, err := h.service.FindPage(page, size, keyword, tagType, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.PageSuccess(c, tags, total, page, size)
}

func (h *GameTagHandler) GetAll(c *gin.Context) {
	tagType, _ := strconv.Atoi(c.DefaultPostForm("type", "1"))
	tags, err := h.service.FindAllByType(tagType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, tags)
}

func (h *GameTagHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	tag, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "标签不存在")
		return
	}

	response.Success(c, tag)
}

func (h *GameTagHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	var tag game.GameTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if err := h.service.Update(uint(id), &tag); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *GameTagHandler) Delete(c *gin.Context) {
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

type GameVariableHandler struct {
	service *gameSvc.GameVariableService
}

func NewGameVariableHandler() *GameVariableHandler {
	return &GameVariableHandler{
		service: gameSvc.NewGameVariableService(),
	}
}

func (h *GameVariableHandler) Create(c *gin.Context) {
	var v game.GameVariable
	if err := c.ShouldBindJSON(&v); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if v.Name == "" || v.Key == "" {
		response.Error(c, http.StatusBadRequest, "变量名称和Key不能为空")
		return
	}

	if err := h.service.Create(&v); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, v)
}

func (h *GameVariableHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))

	vars, total, err := h.service.FindPage(page, size, keyword, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.PageSuccess(c, vars, total, page, size)
}

func (h *GameVariableHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	v, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "变量不存在")
		return
	}

	response.Success(c, v)
}

func (h *GameVariableHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	var v game.GameVariable
	if err := c.ShouldBindJSON(&v); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if err := h.service.Update(uint(id), &v); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *GameVariableHandler) Delete(c *gin.Context) {
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

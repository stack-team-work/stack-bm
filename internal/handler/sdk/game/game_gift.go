package game

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sdk/game"
	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"
	"github.com/gin-gonic/gin"
)

type GameGiftHandler struct{ service *gameSvc.GameGiftService }

func NewGameGiftHandler() *GameGiftHandler {
	return &GameGiftHandler{service: gameSvc.NewGameGiftService()}
}

func (h *GameGiftHandler) Create(c *gin.Context) {
	var g game.GameGift
	if err := c.ShouldBindJSON(&g); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if g.Name == "" {
		response.Error(c, http.StatusBadRequest, "礼包名称不能为空")
		return
	}
	if err := h.service.Create(&g); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, g)
}

func (h *GameGiftHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	list, total, err := h.service.FindPage(page, size, keyword, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *GameGiftHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *GameGiftHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	g, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "礼包不存在")
		return
	}
	response.Success(c, g)
}

func (h *GameGiftHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var g game.GameGift
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

func (h *GameGiftHandler) Delete(c *gin.Context) {
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

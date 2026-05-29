package game

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sdk/game"
	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"
	"github.com/gin-gonic/gin"
)

type GameGiftCodeHandler struct{ service *gameSvc.GameGiftCodeService }

func NewGameGiftCodeHandler() *GameGiftCodeHandler {
	return &GameGiftCodeHandler{service: gameSvc.NewGameGiftCodeService()}
}

func (h *GameGiftCodeHandler) Create(c *gin.Context) {
	var cg game.GameGiftCode
	if err := c.ShouldBindJSON(&cg); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if cg.Code == "" || cg.GiftID == 0 {
		response.Error(c, http.StatusBadRequest, "激活码和礼包ID不能为空")
		return
	}
	if err := h.service.Create(&cg); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, cg)
}

func (h *GameGiftCodeHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	giftID, _ := strconv.Atoi(c.DefaultPostForm("gift_id", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, status, giftID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *GameGiftCodeHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	cg, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "礼包码不存在")
		return
	}
	response.Success(c, cg)
}

func (h *GameGiftCodeHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var cg game.GameGiftCode
	if err := c.ShouldBindJSON(&cg); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &cg); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *GameGiftCodeHandler) Delete(c *gin.Context) {
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

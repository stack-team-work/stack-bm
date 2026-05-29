package game

import (
	"net/http"
	"strconv"

	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"
	"github.com/gin-gonic/gin"
)

type GameGiftUserCodeHandler struct{ service *gameSvc.GameGiftUserCodeService }

func NewGameGiftUserCodeHandler() *GameGiftUserCodeHandler {
	return &GameGiftUserCodeHandler{service: gameSvc.NewGameGiftUserCodeService()}
}

func (h *GameGiftUserCodeHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	giftID, _ := strconv.Atoi(c.DefaultPostForm("gift_id", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, giftID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

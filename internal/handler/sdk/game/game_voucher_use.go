package game

import (
	"net/http"
	"strconv"

	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"
	"github.com/gin-gonic/gin"
)

type GameVoucherUseHandler struct{ service *gameSvc.GameVoucherUseService }

func NewGameVoucherUseHandler() *GameVoucherUseHandler {
	return &GameVoucherUseHandler{service: gameSvc.NewGameVoucherUseService()}
}

func (h *GameVoucherUseHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	voucherID, _ := strconv.Atoi(c.DefaultPostForm("voucher_id", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, voucherID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

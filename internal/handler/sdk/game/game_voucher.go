package game

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sdk/game"
	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"
	"github.com/gin-gonic/gin"
)

type GameVoucherHandler struct{ service *gameSvc.GameVoucherService }

func NewGameVoucherHandler() *GameVoucherHandler {
	return &GameVoucherHandler{service: gameSvc.NewGameVoucherService()}
}

func (h *GameVoucherHandler) Create(c *gin.Context) {
	var v game.GameVoucher
	if err := c.ShouldBindJSON(&v); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if v.Name == "" {
		response.Error(c, http.StatusBadRequest, "代金券名称不能为空")
		return
	}
	if v.AdminID == 0 {
		if uid, ok := c.Get("user_id"); ok {
			v.AdminID = int(uid.(uint))
		}
	}
	if err := h.service.Create(&v); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, v)
}

func (h *GameVoucherHandler) GetList(c *gin.Context) {
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

func (h *GameVoucherHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *GameVoucherHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	v, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "代金券不存在")
		return
	}
	response.Success(c, v)
}

func (h *GameVoucherHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var v game.GameVoucher
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

func (h *GameVoucherHandler) Delete(c *gin.Context) {
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

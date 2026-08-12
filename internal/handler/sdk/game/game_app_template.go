package game

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sdk/game"
	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type GameAppTemplateHandler struct {
	service *gameSvc.GameAppTemplateService
}

func NewGameAppTemplateHandler() *GameAppTemplateHandler {
	return &GameAppTemplateHandler{service: gameSvc.NewGameAppTemplateService()}
}

func (h *GameAppTemplateHandler) Create(c *gin.Context) {
	var t game.GameAppTemplate
	if err := c.ShouldBindJSON(&t); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if t.Name == "" {
		response.Error(c, http.StatusBadRequest, "模板名称不能为空")
		return
	}
	if t.AdminID == 0 {
		if uid, ok := c.Get("user_id"); ok {
			t.AdminID = int(uid.(uint))
		}
	}
	if err := h.service.Create(&t); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, t)
}

func (h *GameAppTemplateHandler) GetList(c *gin.Context) {
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

func (h *GameAppTemplateHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *GameAppTemplateHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	t, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "SDK模板不存在")
		return
	}
	response.Success(c, t)
}

func (h *GameAppTemplateHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var t game.GameAppTemplate
	if err := c.ShouldBindJSON(&t); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &t); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *GameAppTemplateHandler) Delete(c *gin.Context) {
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

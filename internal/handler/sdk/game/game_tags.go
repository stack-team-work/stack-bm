package game

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sdk/game"
	gameSvc "stack-bm/internal/service/sdk/game"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type GameTagHandler struct {
	service *gameSvc.GameTagService
}

func NewGameTagHandler() *GameTagHandler { return &GameTagHandler{service: gameSvc.NewGameTagService()} }

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

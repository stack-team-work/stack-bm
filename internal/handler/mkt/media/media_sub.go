package media

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/mkt/media"
	mediaSvc "stack-bm/internal/service/mkt/media"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type MediaSubHandler struct {
	service *mediaSvc.MediaSubService
}

func NewMediaSubHandler() *MediaSubHandler { return &MediaSubHandler{service: mediaSvc.NewMediaSubService()} }

func (h *MediaSubHandler) Create(c *gin.Context) {
	var s media.MediaSub
	if err := c.ShouldBindJSON(&s); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if s.Name == "" || s.MediaID == 0 {
		response.Error(c, http.StatusBadRequest, "名称和媒体渠道不能为空")
		return
	}
	if err := h.service.Create(&s); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, s)
}

func (h *MediaSubHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	mediaID, _ := strconv.Atoi(c.DefaultPostForm("media_id", "0"))
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	list, total, err := h.service.FindPage(page, size, keyword, mediaID, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *MediaSubHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *MediaSubHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	s, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "媒体子渠道不存在")
		return
	}
	response.Success(c, s)
}

func (h *MediaSubHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var s media.MediaSub
	if err := c.ShouldBindJSON(&s); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &s); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *MediaSubHandler) Delete(c *gin.Context) {
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

package media

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/mkt/media"
	mediaSvc "stack-bm/internal/service/mkt/media"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type MediaApplicationHandler struct {
	service *mediaSvc.MediaApplicationService
}

func NewMediaApplicationHandler() *MediaApplicationHandler {
	return &MediaApplicationHandler{service: mediaSvc.NewMediaApplicationService()}
}

func (h *MediaApplicationHandler) Create(c *gin.Context) {
	var m media.MediaApplication
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if m.Name == "" || m.MediaID == 0 || m.AppID == "" {
		response.Error(c, http.StatusBadRequest, "名称、媒体和AppID不能为空")
		return
	}
	if m.AdminID == 0 {
		if uid, ok := c.Get("user_id"); ok {
			m.AdminID = int(uid.(uint))
		}
	}
	if err := h.service.Create(&m); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, m)
}

func (h *MediaApplicationHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	mediaID, _ := strconv.Atoi(c.DefaultPostForm("media_id", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, status, mediaID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *MediaApplicationHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *MediaApplicationHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	m, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "mkt应用不存在")
		return
	}
	response.Success(c, m)
}

func (h *MediaApplicationHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var m media.MediaApplication
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &m); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *MediaApplicationHandler) Delete(c *gin.Context) {
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

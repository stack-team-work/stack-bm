package media

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/mkt/media"
	mediaSvc "stack-bm/internal/service/mkt/media"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type MediaManagerHandler struct {
	service *mediaSvc.MediaManagerService
}

func NewMediaManagerHandler() *MediaManagerHandler {
	return &MediaManagerHandler{service: mediaSvc.NewMediaManagerService()}
}

func (h *MediaManagerHandler) Create(c *gin.Context) {
	var m media.MediaManager
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if m.Name == "" || m.MediaID == 0 || m.ApplicationID == 0 || m.Account == "" || m.AccountID == "" {
		response.Error(c, http.StatusBadRequest, "必填字段不能为空")
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

func (h *MediaManagerHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	mediaID, _ := strconv.Atoi(c.DefaultPostForm("media_id", "0"))
	applicationID, _ := strconv.Atoi(c.DefaultPostForm("application_id", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, status, mediaID, applicationID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *MediaManagerHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *MediaManagerHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	m, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "mkt管家不存在")
		return
	}
	response.Success(c, m)
}

func (h *MediaManagerHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var m media.MediaManager
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

func (h *MediaManagerHandler) Delete(c *gin.Context) {
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

// Oauth 生成管家授权跳转 URL
func (h *MediaManagerHandler) Oauth(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.ID == 0 {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	url, err := h.service.Oauth(req.ID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"id": req.ID, "url": url})
}

// SyncAdvertiser 同步管家广告主列表
func (h *MediaManagerHandler) SyncAdvertiser(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.ID == 0 {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.service.SyncAdvertiser(req.ID); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, gin.H{"status": true})
}

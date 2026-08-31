package media

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/mkt/media"
	mediaSvc "stack-bm/internal/service/mkt/media"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type MediaDepHandler struct {
	service *mediaSvc.MediaDepService
}

func NewMediaDepHandler() *MediaDepHandler {
	return &MediaDepHandler{service: mediaSvc.NewMediaDepService()}
}

func (h *MediaDepHandler) Create(c *gin.Context) {
	var m media.MediaDep
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if m.Name == "" {
		response.Error(c, http.StatusBadRequest, "名称不能为空")
		return
	}
	if err := h.service.Create(&m); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, m)
}

func (h *MediaDepHandler) GetList(c *gin.Context) {
	var req struct {
		Page    int    `json:"page"`
		Size    int    `json:"size"`
		Keyword string `json:"keyword"`
		Status  *int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 10
	}
	status := -1
	if req.Status != nil {
		status = *req.Status
	}
	list, total, err := h.service.FindPage(req.Page, req.Size, req.Keyword, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, req.Page, req.Size)
}

func (h *MediaDepHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *MediaDepHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	m, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "账户部门不存在")
		return
	}
	response.Success(c, m)
}

func (h *MediaDepHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var m media.MediaDep
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

func (h *MediaDepHandler) Delete(c *gin.Context) {
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

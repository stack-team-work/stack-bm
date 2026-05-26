package sys

import (
	"net/http"
	"strconv"

	bmSysSvc "stack-bm/internal/service/bm/sys"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysLogHandler struct {
	service *bmSysSvc.SysLogService
}

func NewSysLogHandler() *SysLogHandler { return &SysLogHandler{service: bmSysSvc.NewSysLogService()} }

func (h *SysLogHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	level := c.DefaultPostForm("level", "")
	logs, total, err := h.service.FindPage(page, size, keyword, level)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, logs, total, page, size)
}

func (h *SysLogHandler) ClearAll(c *gin.Context) {
	if err := h.service.ClearAll(); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

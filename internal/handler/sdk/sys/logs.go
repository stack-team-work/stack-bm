package sys

import (
	"net/http"
	"strconv"

	sdkSysSvc "stack-bm/internal/service/sdk/sys"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysLogHandler struct {
	service *sdkSysSvc.SysLogService
}

func NewSysLogHandler() *SysLogHandler {
	return &SysLogHandler{service: sdkSysSvc.NewSysLogService()}
}

func (h *SysLogHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	level, _ := strconv.Atoi(c.DefaultPostForm("level", "0"))
	logType, _ := strconv.Atoi(c.DefaultPostForm("type", "0"))

	logs, total, err := h.service.FindPage(page, size, keyword, level, logType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, logs, total, page, size)
}

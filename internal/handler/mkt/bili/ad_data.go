package bili

import (
	"net/http"

	biliSvc "stack-bm/internal/service/mkt/bili"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// AdDataHandler 广告数据
type AdDataHandler struct {
	service *biliSvc.AdDataService
}

func NewAdDataHandler() *AdDataHandler {
	return &AdDataHandler{service: biliSvc.NewAdDataService()}
}

// List 广告数据列表（按层级）
func (h *AdDataHandler) List(c *gin.Context) {
	level := c.Param("level")
	page, _ := strconvAtoi(c.DefaultPostForm("page", "1"))
	size, _ := strconvAtoi(c.DefaultPostForm("size", "10"))
	params := map[string]interface{}{}
	list, total, err := h.service.List(level, page, size, params)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}
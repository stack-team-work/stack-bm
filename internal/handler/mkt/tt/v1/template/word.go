package template

import (
	"net/http"

	ttTplSvc "stack-bm/internal/service/mkt/tt/v1/template"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// WordHandler 头条词包
type WordHandler struct {
	service *ttTplSvc.WordService
}

func NewWordHandler() *WordHandler {
	return &WordHandler{service: ttTplSvc.NewWordService()}
}

// Register 注册词包路由（挂载于 /tt/v1 组）
func (h *WordHandler) Register(rg gin.IRouter) {
	rg.POST("/template/word/list", h.GetList)
}

func (h *WordHandler) GetList(c *gin.Context) {
	list, err := h.service.ListAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

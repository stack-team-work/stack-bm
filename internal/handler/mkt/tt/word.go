package tt

import (
	"net/http"

	ttSvc "stack-bm/internal/service/mkt/tt"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type WordListHandler struct {
	service *ttSvc.WordListService
}

func NewWordListHandler() *WordListHandler {
	return &WordListHandler{service: ttSvc.NewWordListService()}
}

func (h *WordListHandler) GetList(c *gin.Context) {
	list, err := h.service.ListAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

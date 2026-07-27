package handler

import (
	"net/http"

	"stack-bm/pkg/dict"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type DictHandler struct{}

func NewDictHandler() *DictHandler {
	return &DictHandler{}
}

func (h *DictHandler) GetAll(c *gin.Context) {
	response.Success(c, dict.Dict)
}

func (h *DictHandler) GetByKey(c *gin.Context) {
	key := c.Param("key")
	val, ok := dict.Dict[key]
	if !ok {
		response.Error(c, http.StatusNotFound, "字典key不存在")
		return
	}
	response.Success(c, val)
}

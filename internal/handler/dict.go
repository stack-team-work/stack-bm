package handler

import (
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

package tt

import (
	"net/http"

	ttModel "stack-bm/internal/model/mkt/tt"
	ttSvc "stack-bm/internal/service/mkt/tt"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdTemplateHandler struct {
	service *ttSvc.AdTemplateService
}

func NewAdTemplateHandler() *AdTemplateHandler {
	return &AdTemplateHandler{service: ttSvc.NewAdTemplateService()}
}

func (h *AdTemplateHandler) Create(c *gin.Context) {
	var doc ttModel.AdTemplate
	if err := c.ShouldBindJSON(&doc); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Create(&doc, adminID(c)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, doc)
}

func (h *AdTemplateHandler) GetList(c *gin.Context) {
	page, _ := strconvAtoi(c.DefaultPostForm("page", "1"))
	size, _ := strconvAtoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	list, total, err := h.service.FindPage(page, size, keyword)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *AdTemplateHandler) GetByID(c *gin.Context) {
	doc, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusNotFound, "广告模板不存在")
		return
	}
	response.Success(c, doc)
}

func (h *AdTemplateHandler) Update(c *gin.Context) {
	var doc ttModel.AdTemplate
	if err := c.ShouldBindJSON(&doc); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(c.Param("id"), &doc); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AdTemplateHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AdTemplateHandler) Copy(c *gin.Context) {
	var req struct {
		ID           string `json:"id"`
		TemplateName string `json:"template_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Copy(req.ID, req.TemplateName, adminID(c)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

package bili

import (
	"net/http"

	biliModel "stack-bm/internal/model/mkt/bili"
	biliSvc "stack-bm/internal/service/mkt/bili"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type AudienceTemplateHandler struct {
	service *biliSvc.AudienceTemplateService
}

func NewAudienceTemplateHandler() *AudienceTemplateHandler {
	return &AudienceTemplateHandler{service: biliSvc.NewAudienceTemplateService()}
}

func (h *AudienceTemplateHandler) Create(c *gin.Context) {
	var doc biliModel.AudienceTemplate
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

func (h *AudienceTemplateHandler) GetList(c *gin.Context) {
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

func (h *AudienceTemplateHandler) GetByID(c *gin.Context) {
	doc, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusNotFound, "定向包模板不存在")
		return
	}
	response.Success(c, doc)
}

func (h *AudienceTemplateHandler) Update(c *gin.Context) {
	var doc biliModel.AudienceTemplate
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

func (h *AudienceTemplateHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AudienceTemplateHandler) Copy(c *gin.Context) {
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

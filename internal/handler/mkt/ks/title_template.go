package ks

import (
	"net/http"

	ksModel "stack-bm/internal/model/mkt/ks"
	ksSvc "stack-bm/internal/service/mkt/ks"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type TitleTemplateHandler struct {
	service *ksSvc.TitleTemplateService
}

func NewTitleTemplateHandler() *TitleTemplateHandler {
	return &TitleTemplateHandler{service: ksSvc.NewTitleTemplateService()}
}

func (h *TitleTemplateHandler) Create(c *gin.Context) {
	var doc ksModel.TitleTemplate
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

func (h *TitleTemplateHandler) GetList(c *gin.Context) {
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

func (h *TitleTemplateHandler) GetByID(c *gin.Context) {
	doc, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusNotFound, "标题包模板不存在")
		return
	}
	response.Success(c, doc)
}

func (h *TitleTemplateHandler) Update(c *gin.Context) {
	var doc ksModel.TitleTemplate
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

func (h *TitleTemplateHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *TitleTemplateHandler) Copy(c *gin.Context) {
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

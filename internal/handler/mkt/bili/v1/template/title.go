package template

import (
	"net/http"

	biliModel "stack-bm/internal/model/mkt/bili"
	biliTplSvc "stack-bm/internal/service/mkt/bili/v1/template"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// TitleHandler B站标题包模板
type TitleHandler struct {
	service *biliTplSvc.TitleService
}

func NewTitleHandler() *TitleHandler {
	return &TitleHandler{service: biliTplSvc.NewTitleService()}
}

// Register 注册标题包模板路由（挂载于 /bili/v1 组）
func (h *TitleHandler) Register(rg gin.IRouter) {
	prefix := "/template/title/"
	rg.POST(prefix+"create", h.Create)
	rg.POST(prefix+"list", h.GetList)
	rg.POST(prefix+"detail/:id", h.GetByID)
	rg.POST(prefix+"update/:id", h.Update)
	rg.POST(prefix+"delete/:id", h.Delete)
	rg.POST(prefix+"copy", h.Copy)
}

func (h *TitleHandler) Create(c *gin.Context) {
	var doc biliModel.TitleTemplate
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

func (h *TitleHandler) GetList(c *gin.Context) {
	var req listRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	req.normalize()
	list, total, err := h.service.FindPage(req.Page, req.Size, req.Keyword)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, req.Page, req.Size)
}

func (h *TitleHandler) GetByID(c *gin.Context) {
	doc, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusNotFound, "标题包模板不存在")
		return
	}
	response.Success(c, doc)
}

func (h *TitleHandler) Update(c *gin.Context) {
	var doc biliModel.TitleTemplate
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

func (h *TitleHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *TitleHandler) Copy(c *gin.Context) {
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

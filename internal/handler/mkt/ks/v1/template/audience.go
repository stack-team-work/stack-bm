package template

import (
	"net/http"

	ksModel "stack-bm/internal/model/mkt/ks"
	ksTplSvc "stack-bm/internal/service/mkt/ks/v1/template"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

// AudienceHandler B站定向包模板
type AudienceHandler struct {
	service *ksTplSvc.AudienceService
}

func NewAudienceHandler() *AudienceHandler {
	return &AudienceHandler{service: ksTplSvc.NewAudienceService()}
}

// Register 注册定向包模板路由（挂载于 /bili/v1 组）
func (h *AudienceHandler) Register(rg gin.IRouter) {
	prefix := "/template/audience/"
	rg.POST(prefix+"create", h.Create)
	rg.POST(prefix+"list", h.GetList)
	rg.POST(prefix+"detail/:id", h.GetByID)
	rg.POST(prefix+"update/:id", h.Update)
	rg.POST(prefix+"delete/:id", h.Delete)
	rg.POST(prefix+"copy", h.Copy)
}

func (h *AudienceHandler) Create(c *gin.Context) {
	var doc ksModel.AudienceTemplate
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

func (h *AudienceHandler) GetList(c *gin.Context) {
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

func (h *AudienceHandler) GetByID(c *gin.Context) {
	doc, err := h.service.FindByID(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusNotFound, "定向包模板不存在")
		return
	}
	response.Success(c, doc)
}

func (h *AudienceHandler) Update(c *gin.Context) {
	var doc ksModel.AudienceTemplate
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

func (h *AudienceHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Param("id")); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *AudienceHandler) Copy(c *gin.Context) {
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

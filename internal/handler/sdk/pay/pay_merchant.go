package pay

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sdk/pay"
	paySvc "stack-bm/internal/service/sdk/pay"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type PayMerchantHandler struct {
	service *paySvc.PayMerchantService
}

func NewPayMerchantHandler() *PayMerchantHandler {
	return &PayMerchantHandler{service: paySvc.NewPayMerchantService()}
}

func (h *PayMerchantHandler) Create(c *gin.Context) {
	var p pay.PayMerchant
	if err := c.ShouldBindJSON(&p); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if p.Name == "" || p.ShowName == "" {
		response.Error(c, http.StatusBadRequest, "名称不能为空")
		return
	}
	if err := h.service.Create(&p); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, p)
}

func (h *PayMerchantHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	payType, _ := strconv.Atoi(c.DefaultPostForm("type", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, status, payType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *PayMerchantHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *PayMerchantHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	p, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "支付商户不存在")
		return
	}
	response.Success(c, p)
}

func (h *PayMerchantHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var p pay.PayMerchant
	if err := c.ShouldBindJSON(&p); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &p); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *PayMerchantHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

package sys

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/bm/sys"
	bmSysSvc "stack-bm/internal/service/bm/sys"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysTagHandler struct {
	service *bmSysSvc.SysTagService
}

func NewSysTagHandler() *SysTagHandler {
	return &SysTagHandler{service: bmSysSvc.NewSysTagService()}
}

func (h *SysTagHandler) Create(c *gin.Context) {
	var t sys.SysTag
	if err := c.ShouldBindJSON(&t); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if t.AdminID == 0 {
		if uid, ok := c.Get("user_id"); ok {
			t.AdminID = int(uid.(uint))
		}
	}
	if err := h.service.Create(&t); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, t)
}

func (h *SysTagHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	tagType, _ := strconv.Atoi(c.DefaultPostForm("type", "0"))
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	list, total, err := h.service.FindPage(page, size, keyword, tagType, status)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *SysTagHandler) GetAll(c *gin.Context) {
	tagType, _ := strconv.Atoi(c.DefaultPostForm("type", "0"))
	options, err := h.service.FindOptions(tagType)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, options)
}

func (h *SysTagHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	t, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "标签不存在")
		return
	}
	response.Success(c, t)
}

func (h *SysTagHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var t sys.SysTag
	if err := c.ShouldBindJSON(&t); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &t); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

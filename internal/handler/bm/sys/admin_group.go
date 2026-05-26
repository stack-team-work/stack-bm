package sys

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/bm/sys"
	bmSysSvc "stack-bm/internal/service/bm/sys"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysAdminGroupHandler struct {
	service *bmSysSvc.SysAdminGroupService
}

func NewSysAdminGroupHandler() *SysAdminGroupHandler {
	return &SysAdminGroupHandler{service: bmSysSvc.NewSysAdminGroupService()}
}

func (h *SysAdminGroupHandler) Create(c *gin.Context) {
	var group sys.SysAdminGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if group.Name == "" {
		response.Error(c, http.StatusBadRequest, "分组名称不能为空")
		return
	}
	if err := h.service.Create(&group); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, group)
}

func (h *SysAdminGroupHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	groups, total, err := h.service.FindPage(page, size, keyword)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, groups, total, page, size)
}

func (h *SysAdminGroupHandler) GetAll(c *gin.Context) {
	groups, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, groups)
}

func (h *SysAdminGroupHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	group, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "分组不存在")
		return
	}
	response.Success(c, group)
}

func (h *SysAdminGroupHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var group sys.SysAdminGroup
	if err := c.ShouldBindJSON(&group); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &group); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SysAdminGroupHandler) Delete(c *gin.Context) {
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

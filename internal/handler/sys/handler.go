package sys

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/sys"
	sysSvc "stack-bm/internal/service/sys"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysAdminHandler struct {
	service *sysSvc.SysAdminService
}

func NewSysAdminHandler() *SysAdminHandler {
	return &SysAdminHandler{
		service: sysSvc.NewSysAdminService(),
	}
}

func (h *SysAdminHandler) Create(c *gin.Context) {
	var admin sys.SysAdmin
	if err := c.ShouldBindJSON(&admin); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if admin.Username == "" || admin.Password == "" {
		response.Error(c, http.StatusBadRequest, "用户名和密码不能为空")
		return
	}

	if err := h.service.Create(&admin); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, admin)
}

func (h *SysAdminHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	groupID, _ := strconv.Atoi(c.DefaultPostForm("group_id", "0"))

	admins, total, err := h.service.FindPage(page, size, keyword, groupID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.PageSuccess(c, admins, total, page, size)
}

func (h *SysAdminHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	admin, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "管理员不存在")
		return
	}

	response.Success(c, admin)
}

func (h *SysAdminHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}

	var admin sys.SysAdmin
	if err := c.ShouldBindJSON(&admin); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if err := h.service.Update(uint(id), &admin); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, nil)
}

func (h *SysAdminHandler) Delete(c *gin.Context) {
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

type SysAdminGroupHandler struct {
	service *sysSvc.SysAdminGroupService
}

func NewSysAdminGroupHandler() *SysAdminGroupHandler {
	return &SysAdminGroupHandler{
		service: sysSvc.NewSysAdminGroupService(),
	}
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

type SysLogHandler struct {
	service *sysSvc.SysLogService
}

func NewSysLogHandler() *SysLogHandler {
	return &SysLogHandler{service: sysSvc.NewSysLogService()}
}

func (h *SysLogHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	level := c.DefaultPostForm("level", "")

	logs, total, err := h.service.FindPage(page, size, keyword, level)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, logs, total, page, size)
}

func (h *SysLogHandler) ClearAll(c *gin.Context) {
	if err := h.service.ClearAll(); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

type SysMenuHandler struct {
	service *sysSvc.SysMenuService
}

func NewSysMenuHandler() *SysMenuHandler {
	return &SysMenuHandler{service: sysSvc.NewSysMenuService()}
}

func (h *SysMenuHandler) Create(c *gin.Context) {
	var m sys.SysMenu
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if m.Name == "" || m.Path == "" {
		response.Error(c, http.StatusBadRequest, "菜单名称和路径不能为空")
		return
	}
	if err := h.service.Create(&m); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, m)
}

func (h *SysMenuHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "50"))
	menus, total, err := h.service.FindPage(page, size)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, menus, total, page, size)
}

func (h *SysMenuHandler) GetAll(c *gin.Context) {
	menus, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, menus)
}

func (h *SysMenuHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	m, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "菜单不存在")
		return
	}
	response.Success(c, m)
}

func (h *SysMenuHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var m sys.SysMenu
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if err := h.service.Update(uint(id), &m); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

func (h *SysMenuHandler) Delete(c *gin.Context) {
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

package sys

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/bm/sys"
	bmSysSvc "stack-bm/internal/service/bm/sys"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysMenuHandler struct {
	service *bmSysSvc.SysMenuService
}

func NewSysMenuHandler() *SysMenuHandler { return &SysMenuHandler{service: bmSysSvc.NewSysMenuService()} }

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

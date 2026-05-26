package sys

import (
	"net/http"
	"strconv"

	"stack-bm/internal/model/bm/sys"
	bmSysSvc "stack-bm/internal/service/bm/sys"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysAdminHandler struct {
	service *bmSysSvc.SysAdminService
}

func NewSysAdminHandler() *SysAdminHandler { return &SysAdminHandler{service: bmSysSvc.NewSysAdminService()} }

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

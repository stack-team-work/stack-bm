package sys

import (
	"net/http"
	"strconv"
	"stack-bm/internal/model/bm/sys"
	bmSysSvc "stack-bm/internal/service/bm/sys"
	"stack-bm/pkg/response"
	"github.com/gin-gonic/gin"
)

type FeishuUserHandler struct{ service *bmSysSvc.FeishuUserService }

func NewFeishuUserHandler() *FeishuUserHandler {
	return &FeishuUserHandler{service: bmSysSvc.NewFeishuUserService()}
}

func (h *FeishuUserHandler) Create(c *gin.Context) {
	var m sys.FeishuUser
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if m.FeishuUserID == "" {
		response.Error(c, http.StatusBadRequest, "飞书用户ID不能为空")
		return
	}
	if err := h.service.Create(&m); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, m)
}

func (h *FeishuUserHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	adminID, _ := strconv.Atoi(c.DefaultPostForm("admin_id", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, status, adminID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *FeishuUserHandler) GetAll(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	adminSvc := bmSysSvc.NewSysAdminService()
	type Item struct {
		ID           uint   `json:"id"`
		AdminID      int    `json:"admin_id"`
		FeishuUserID string `json:"feishu_user_id"`
		AdminName    string `json:"admin_name"`
		Status       int    `json:"status"`
	}
	list := make([]Item, 0, len(users))
	for _, u := range users {
		item := Item{ID: u.ID, AdminID: u.AdminID, FeishuUserID: u.FeishuUserID, Status: u.Status}
		if admin, err := adminSvc.FindByID(uint(u.AdminID)); err == nil {
			item.AdminName = admin.Name
		}
		list = append(list, item)
	}
	response.Success(c, list)
}

func (h *FeishuUserHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	m, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "飞书用户不存在")
		return
	}
	response.Success(c, m)
}

func (h *FeishuUserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var m sys.FeishuUser
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

func (h *FeishuUserHandler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var body struct{ Status int `json:"status"` }
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	if err := h.service.UpdateStatus(uint(id), body.Status); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, nil)
}

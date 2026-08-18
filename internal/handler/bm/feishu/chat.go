package feishu

import (
	"net/http"
	"strconv"
	"stack-bm/internal/model/bm/feishu"
	feishuSvc "stack-bm/internal/service/bm/feishu"
	"stack-bm/pkg/response"
	"github.com/gin-gonic/gin"
)

type FeishuChatHandler struct{ service *feishuSvc.FeishuChatService }

func NewFeishuChatHandler() *FeishuChatHandler {
	return &FeishuChatHandler{service: feishuSvc.NewFeishuChatService()}
}

func (h *FeishuChatHandler) Create(c *gin.Context) {
	var m feishu.FeishuChat
	if err := c.ShouldBindJSON(&m); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}
	if m.ChatID == "" || m.CallAction == "" {
		response.Error(c, http.StatusBadRequest, "ChatID和CallAction不能为空")
		return
	}
	if m.AdminID == 0 {
		if uid, ok := c.Get("user_id"); ok {
			m.AdminID = int(uid.(uint))
		}
	}
	if err := h.service.Create(&m); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, m)
}

func (h *FeishuChatHandler) GetList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "10"))
	keyword := c.DefaultPostForm("keyword", "")
	status, _ := strconv.Atoi(c.DefaultPostForm("status", "-1"))
	feishuAppID, _ := strconv.Atoi(c.DefaultPostForm("feishu_app_id", "0"))
	list, total, err := h.service.FindPage(page, size, keyword, status, feishuAppID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.PageSuccess(c, list, total, page, size)
}

func (h *FeishuChatHandler) GetAll(c *gin.Context) {
	list, err := h.service.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, list)
}

func (h *FeishuChatHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	m, err := h.service.FindByID(uint(id))
	if err != nil {
		response.Error(c, http.StatusNotFound, "飞书聊天不存在")
		return
	}
	response.Success(c, m)
}

func (h *FeishuChatHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误")
		return
	}
	var m feishu.FeishuChat
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

func (h *FeishuChatHandler) UpdateStatus(c *gin.Context) {
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

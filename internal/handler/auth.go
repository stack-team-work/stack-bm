package handler

import (
	"net/http"

	"stack-bm/internal/service"
	"stack-bm/pkg/jwt"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: service.NewAuthService(),
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	admin, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := jwt.GenerateToken(admin.ID, admin.Username)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "生成token失败")
		return
	}

	response.Success(c, gin.H{
		"token": token,
		"admin": gin.H{
			"id":       admin.ID,
			"username": admin.Username,
			"name":     admin.Name,
			"phone":    admin.Phone,
			"group_id": admin.GroupID,
		},
	})
}

func (h *AuthHandler) GetUserInfo(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "未登录")
		return
	}

	userClaims := claims.(*jwt.Claims)

	response.Success(c, gin.H{
		"id":       userClaims.UserID,
		"username": userClaims.Username,
	})
}

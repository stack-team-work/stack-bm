package handler

import (
	"net/http"

	"stack-bm/internal/database"
	"stack-bm/internal/service"
	"stack-bm/pkg/captcha"
	"stack-bm/pkg/constants"
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
	CaptchaID string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

func (h *AuthHandler) Captcha(c *gin.Context) {
	id, question := captcha.Generate()
	response.Success(c, gin.H{"captcha_id": id, "question": question})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "参数错误: "+err.Error())
		return
	}

	if !captcha.Verify(req.CaptchaID, req.Captcha) {
		writeLog(constants.BM_LOG_LEVEL_ERROR, "/api/login", req.Username, c.ClientIP(), "验证码错误")
		response.Error(c, http.StatusBadRequest, "验证码错误")
		return
	}

	admin, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		writeLog(constants.BM_LOG_LEVEL_ERROR, "/api/login", req.Username, c.ClientIP(), err.Error())
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := jwt.GenerateToken(admin.ID, admin.Username)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "生成token失败")
		return
	}

	writeLog(constants.BM_LOG_LEVEL_INFO, "/api/login", req.Username, c.ClientIP(), "用户登录成功")
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

func writeLog(level int, path, username, ip, desc string) {
	if database.DBBM != nil {
		database.DBBM.Exec(`INSERT INTO sys_logs (level, path, username, ip, `+"`desc`"+`, created_at, updated_at) VALUES (?, ?, ?, ?, ?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())`,
			level, path, username, ip, desc)
	}
}

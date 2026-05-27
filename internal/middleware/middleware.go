package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"stack-bm/internal/database"
	"stack-bm/pkg/dict"
	"stack-bm/pkg/jwt"
	"stack-bm/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "未提供认证令牌")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "认证令牌格式错误")
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "认证令牌无效或已过期")
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		statusCode := c.Writer.Status()
		if statusCode < 400 {
			return
		}

		username, _ := c.Get("username")
		uname := ""
		if username != nil {
			uname = username.(string)
		}

		desc := string(bodyBytes)
		if len(desc) > 500 {
			desc = desc[:500]
		}

		db := database.DBBM
		if db != nil {
			db.Exec(`INSERT INTO sys_logs (level, path, username, ip, `+"`desc`"+`, created_at, updated_at) VALUES (?, ?, ?, ?, ?, UNIX_TIMESTAMP(), UNIX_TIMESTAMP())`,
				dict.BM_LOG_LEVEL_ERROR, c.Request.URL.Path, uname, c.ClientIP(), desc)
		}

		_ = json.Unmarshal
	}
}

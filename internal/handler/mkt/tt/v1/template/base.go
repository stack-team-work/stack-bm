package template

import (
	"github.com/gin-gonic/gin"
)

// adminID 从上下文读取当前管理员ID
func adminID(c *gin.Context) int {
	if uid, ok := c.Get("user_id"); ok {
		return int(uid.(uint))
	}
	return 0
}

// listRequest 模板列表通用入参（JSON）
type listRequest struct {
	Page    int    `json:"page"`
	Size    int    `json:"size"`
	Keyword string `json:"keyword"`
}

// normalize 兜底分页参数
func (r *listRequest) normalize() {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.Size < 1 {
		r.Size = 10
	}
}

package tc

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func adminID(c *gin.Context) int {
	if uid, ok := c.Get("user_id"); ok {
		return int(uid.(uint))
	}
	return 0
}

func strconvAtoi(s string) (int, error) {
	return strconv.Atoi(s)
}
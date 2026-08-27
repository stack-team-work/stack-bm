package template

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册B站 v1 全部模板路由，挂载于形如 /api/bili/v1 的路由组
func RegisterRoutes(rg gin.IRouter) {
	NewAdHandler().Register(rg)
	NewAudienceHandler().Register(rg)
	NewTitleHandler().Register(rg)
}

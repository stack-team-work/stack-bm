package template

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册头条 v1 全部模板路由（含词包），挂载于形如 /api/tt/v1 的路由组
func RegisterRoutes(rg gin.IRouter) {
	NewAdHandler().Register(rg)
	NewAudienceHandler().Register(rg)
	NewTitleHandler().Register(rg)
	NewWordHandler().Register(rg)
}

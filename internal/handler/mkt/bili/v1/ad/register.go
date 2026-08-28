package ad

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册B站 v1 广告数据全部层级路由，挂载于形如 /api/bili/v1 的路由组
func RegisterRoutes(rg gin.IRouter) {
	g := rg.Group("/ad")
	NewAccountHandler().Register(g)
	NewCampaignHandler().Register(g)
	NewUnitHandler().Register(g)
	NewCreativeHandler().Register(g)
}

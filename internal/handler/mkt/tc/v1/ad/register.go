package ad

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册腾讯 v1 广告数据全部层级路由，挂载于形如 /api/tc/v1 的路由组
// 腾讯V3仅 账户 / 广告组(campaign) 两层
func RegisterRoutes(rg gin.IRouter) {
	g := rg.Group("/ad")
	NewAccountHandler().Register(g)
	NewCampaignHandler().Register(g)
}

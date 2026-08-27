package ad

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册头条 v1 广告数据全部层级路由，挂载于形如 /api/tt/v1 的路由组
// 头条V3仅 账户 / 项目(campaign) 两层
func RegisterRoutes(rg gin.IRouter) {
	NewAccountHandler().Register(rg)
	NewCampaignHandler().Register(rg)
}

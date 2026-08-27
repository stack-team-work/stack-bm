package sync

import (
	"fmt"
)

// Run 单渠道单账户同步入口（供外部 cron 经 runner 调用）
// 头条V3仅支持 项目(campaign) 层级与账户余额占位
func Run(level string, accountID uint) error {
	switch level {
	case "account":
		return SyncBalance(accountID)
	case "campaign":
		return SyncCampaign(accountID)
	default:
		return fmt.Errorf("头条V3不支持同步层级: %s", level)
	}
}

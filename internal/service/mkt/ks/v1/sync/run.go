package sync

import (
	"fmt"
)

// Run 单渠道单账户同步入口（供外部 cron 经 runner 调用）
func Run(level string, accountID uint) error {
	switch level {
	case "account":
		return SyncBalance(accountID)
	case "campaign":
		return SyncCampaign(accountID)
	case "unit":
		return SyncUnit(accountID)
	case "creative":
		return SyncCreative(accountID)
	default:
		return fmt.Errorf("快手不支持同步层级: %s", level)
	}
}

package dict

import "stack-bm/pkg/constants"

type Option struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

var Dict = map[string]interface{}{
	"status": []Option{
		{Label: "启用", Value: constants.STATUS_ENABLE},
		{Label: "禁用", Value: constants.STATUS_DISABLE},
	},
	"game_tag_type": []Option{
		{Label: "风格", Value: constants.GAME_TAG_TYPE_STYLE},
		{Label: "类型", Value: constants.GAME_TAG_TYPE_TYPE},
	},
	"game_app_os": []Option{
		{Label: "安卓", Value: constants.GAME_APP_OS_ANDROID},
		{Label: "iOS", Value: constants.GAME_APP_OS_IOS},
		{Label: "双端", Value: constants.GAME_APP_OS_BOTH},
	},
	"sdk_log_type": []Option{
		{Label: "注册日志", Value: constants.SDK_LOG_TYPE_REGISTER},
		{Label: "登录日志", Value: constants.SDK_LOG_TYPE_LOGIN},
		{Label: "支付日志", Value: constants.SDK_LOG_TYPE_PAY},
	},
	"sdk_log_level": []Option{
		{Label: "info", Value: constants.SDK_LOG_LEVEL_INFO},
		{Label: "warning", Value: constants.SDK_LOG_LEVEL_WARNING},
		{Label: "error", Value: constants.SDK_LOG_LEVEL_ERROR},
	},
	"bm_log_level": []Option{
		{Label: "info", Value: constants.BM_LOG_LEVEL_INFO},
		{Label: "error", Value: constants.BM_LOG_LEVEL_ERROR},
	},
	"menu_type": []Option{
		{Label: "菜单", Value: constants.MENU_TYPE_MENU},
		{Label: "按钮", Value: constants.MENU_TYPE_BUTTON},
	},
	"media_manager_auth_status": []Option{
		{Label: "未授权", Value: constants.MEDIA_MANAGER_AUTH_STATUS_NO},
		{Label: "已授权", Value: constants.MEDIA_MANAGER_AUTH_STATUS_YES},
	},
	"pay_merchant_type": []Option{
		{Label: "微信", Value: constants.PAY_MERCHANT_TYPE_WECHAT},
		{Label: "支付宝", Value: constants.PAY_MERCHANT_TYPE_ALIPAY},
	},
	"game_gift_get_type": []Option{
		{Label: "单次领取", Value: constants.GAME_GIFT_GET_TYPE_ONCE},
		{Label: "每日领取", Value: constants.GAME_GIFT_GET_TYPE_DAILY},
	},
	"game_voucher_use_type": []Option{
		{Label: "玩家角色", Value: constants.GAME_VOUCHER_USE_TYPE_ROLE},
		{Label: "SDK账户", Value: constants.GAME_VOUCHER_USE_TYPE_ACCOUNT},
	},
	"game_gift_type": []Option{
		{Label: "道具", Value: constants.GAME_GIFT_TYPE_PROP},
	},
	"sys_column_report_type": []Option{
		{Label: "投放报表", Value: constants.SYS_COLUMN_REPORT_TYPE_AD},
	},
	"sys_column_indicator_type": []Option{
		{Label: "属性指标", Value: constants.SYS_COLUMN_INDICATOR_ATTR},
		{Label: "媒体指标", Value: constants.SYS_COLUMN_INDICATOR_MEDIA},
		{Label: "BM指标", Value: constants.SYS_COLUMN_INDICATOR_BM},
		{Label: "N日指标", Value: constants.SYS_COLUMN_INDICATOR_NDAY},
	},
	"feishu_chat_type": []Option{
		{Label: "普通机器人", Value: constants.FEISHU_CHAT_TYPE_NORMAL},
		{Label: "应用机器人", Value: constants.FEISHU_CHAT_TYPE_APP},
	},
	"feishu_chat_at_type": []Option{
		{Label: "艾特全部", Value: constants.FEISHU_CHAT_AT_TYPE_ALL},
		{Label: "艾特负责人", Value: constants.FEISHU_CHAT_AT_TYPE_OWNER},
	},
}

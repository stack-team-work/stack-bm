package dict

const (
	STATUS_ENABLE  = 1
	STATUS_DISABLE = 0
)

const (
	GAME_TAG_TYPE_STYLE = 1
	GAME_TAG_TYPE_TYPE  = 2
)

const (
	GAME_APP_OS_ANDROID = 1
	GAME_APP_OS_IOS     = 2
	GAME_APP_OS_BOTH    = 3
)

const (
	SDK_LOG_TYPE_REGISTER = 1
	SDK_LOG_TYPE_LOGIN    = 2
	SDK_LOG_TYPE_PAY      = 3
)

const (
	SDK_LOG_LEVEL_INFO    = 1
	SDK_LOG_LEVEL_WARNING = 2
	SDK_LOG_LEVEL_ERROR   = 3
)

const (
	BM_LOG_LEVEL_INFO  = 1
	BM_LOG_LEVEL_ERROR = 2
)

const (
	MENU_TYPE_MENU   = 1
	MENU_TYPE_BUTTON = 2
)

const (
	MEDIA_MANAGER_AUTH_STATUS_NO  = 0
	MEDIA_MANAGER_AUTH_STATUS_YES = 1
)

const (
	PAY_MERCHANT_TYPE_WECHAT  = 1
	PAY_MERCHANT_TYPE_ALIPAY  = 2
)

type Option struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type OptionStr struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

var Dict = map[string]interface{}{
	"status": []Option{
		{Label: "启用", Value: STATUS_ENABLE},
		{Label: "禁用", Value: STATUS_DISABLE},
	},
	"game_tag_type": []Option{
		{Label: "风格", Value: GAME_TAG_TYPE_STYLE},
		{Label: "类型", Value: GAME_TAG_TYPE_TYPE},
	},
	"game_app_os": []Option{
		{Label: "安卓", Value: GAME_APP_OS_ANDROID},
		{Label: "iOS", Value: GAME_APP_OS_IOS},
		{Label: "双端", Value: GAME_APP_OS_BOTH},
	},
	"sdk_log_type": []Option{
		{Label: "注册日志", Value: SDK_LOG_TYPE_REGISTER},
		{Label: "登录日志", Value: SDK_LOG_TYPE_LOGIN},
		{Label: "支付日志", Value: SDK_LOG_TYPE_PAY},
	},
	"sdk_log_level": []Option{
		{Label: "info", Value: SDK_LOG_LEVEL_INFO},
		{Label: "warning", Value: SDK_LOG_LEVEL_WARNING},
		{Label: "error", Value: SDK_LOG_LEVEL_ERROR},
	},
	"bm_log_level": []Option{
		{Label: "info", Value: BM_LOG_LEVEL_INFO},
		{Label: "error", Value: BM_LOG_LEVEL_ERROR},
	},
	"menu_type": []Option{
		{Label: "菜单", Value: MENU_TYPE_MENU},
		{Label: "按钮", Value: MENU_TYPE_BUTTON},
	},
	"media_manager_auth_status": []Option{
		{Label: "未授权", Value: MEDIA_MANAGER_AUTH_STATUS_NO},
		{Label: "已授权", Value: MEDIA_MANAGER_AUTH_STATUS_YES},
	},
	"pay_merchant_type": []Option{
		{Label: "微信", Value: PAY_MERCHANT_TYPE_WECHAT},
		{Label: "支付宝", Value: PAY_MERCHANT_TYPE_ALIPAY},
	},
}

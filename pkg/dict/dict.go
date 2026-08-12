package dict

import "stack-bm/pkg/constants"

type Option struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type StringOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func biliOpts(in []constants.BiliOption) []Option {
	out := make([]Option, 0, len(in))
	for _, o := range in {
		out = append(out, Option{Label: o.Label, Value: o.Value})
	}
	return out
}

func biliSOpts(in []constants.BiliStringOption) []StringOption {
	out := make([]StringOption, 0, len(in))
	for _, o := range in {
		out = append(out, StringOption{Label: o.Label, Value: o.Value})
	}
	return out
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
	"bili_promotion_purpose":   biliOpts(constants.BiliPromotionPurposeOptions),
	"bili_ad_type":             biliOpts(constants.BiliAdTypeOptions),
	"bili_delivery_type":       biliOpts(constants.BiliDeliveryOptions),
	"bili_budget_type":         biliOpts(constants.BiliBudgetOptions),
	"bili_speed_mode":          biliOpts(constants.BiliSpeedModeOptions),
	"bili_promotion_content":   biliOpts(constants.BiliPromotionContentOptions),
	"bili_mini_game":           biliOpts(constants.BiliMiniGameOptions),
	"bili_unit_date_type":      biliOpts(constants.BiliUnitDateTypeOptions),
	"bili_unit_time_type":      biliOpts(constants.BiliUnitTimeTypeOptions),
	"bili_strategy_type":       biliOpts(constants.BiliStrategyOptions),
	"bili_bid_type":            biliOpts(constants.BiliBidOptions),
	"bili_optimize_gold":       biliOpts(constants.BiliOptimizeGoldOptions),
	"bili_deep_optimize_gold":  biliOpts(constants.BiliDeepOptimizeGoldOptions),
	"bili_deep_gold_type":      biliOpts(constants.BiliDeepGoldOptions),
	"bili_launch_type":         biliOpts(constants.BiliLaunchOptions),
	"bili_network_type":        biliOpts(constants.BiliNetworkOptions),
	"bili_custom_bid_type":     biliSOpts(constants.BiliCustomBidOptions),
	"bili_brand_type":          biliOpts(constants.BiliBrandOptions),
	"bili_area_type":           biliOpts(constants.BiliAreaTypeOptions),
	"bili_installed_type":      biliOpts(constants.BiliInstalledOptions),
	"bili_installed_app_type":  biliOpts(constants.BiliInstalledAppOptions),
	"bili_tag_fuzzy_type":      biliOpts(constants.BiliTagFuzzyOptions),
	"bili_title_word_type":     biliOpts(constants.BiliTitleWordOptions),
	"bili_gender":              biliOpts(constants.BiliGenderOptions),
	"bili_age":                 biliOpts(constants.BiliAgeOptions),
	"bili_network":             biliOpts(constants.BiliNetOptions),
	"bili_phone_price":         biliOpts(constants.BiliPhonePriceOptions),
	"bili_os":                  biliOpts(constants.BiliOsOptions),
	"bili_converted_user":      biliOpts(constants.BiliConvertedUserFilterOptions),
	"ks_market_target":         biliOpts(constants.KsMarketTargetOptions),
	"ks_scene":                 biliOpts(constants.KsSceneOptions),
	"ks_ad_type":               biliOpts(constants.KsAdTypeOptions),
	"ks_auto_manager":          biliOpts(constants.KsAutoManagerOptions),
	"ks_budget_type":           biliOpts(constants.KsBudgetTypeOptions),
	"ks_bid_strategy":          biliOpts(constants.KsBidStrategyOptions),
	"ks_auto_photo_scope":      biliOpts(constants.KsAutoPhotoScopeOptions),
	"ks_creative_unit":         biliOpts(constants.KsCreativeUnitOptions),
	"ks_mini_type":             biliOpts(constants.KsMiniTypeOptions),
	"ks_unit_date_type":        biliOpts(constants.KsUnitDateTypeOptions),
	"ks_unit_time_type":        biliOpts(constants.KsUnitTimeTypeOptions),
	"ks_bid_type":              biliOpts(constants.KsBidTypeOptions),
	"ks_ocpx_action":           biliOpts(constants.KsOcpxActionOptions),
	"ks_deep_conversion":       biliOpts(constants.KsDeepConversionOptions),
	"ks_show_mode":             biliOpts(constants.KsShowModeOptions),
	"ks_bid_way":               biliOpts(constants.KsBidWayOptions),
	"ks_speed_type":            biliOpts(constants.KsSpeedTypeOptions),
	"ks_smart_bid":             biliOpts(constants.KsSmartBidOptions),
	"ks_gender":                biliOpts(constants.KsGenderOptions),
	"ks_age":                   biliOpts(constants.KsAgeOptions),
	"ks_network":               biliOpts(constants.KsNetworkOptions),
	"ks_operators":             biliOpts(constants.KsOperatorsOptions),
	"ks_platform_os":           biliOpts(constants.KsPlatformOsOptions),
	"ks_device_price_type":     biliOpts(constants.KsDevicePriceTypeOptions),
	"ks_device_price":          biliOpts(constants.KsDevicePriceOptions),
	"ks_app_interest_type":     biliOpts(constants.KsAppInterestTypeOptions),
	"ks_filter_converted":      biliOpts(constants.KsFilterConvertedLevelOptions),
	"ks_filter_time":           biliOpts(constants.KsFilterTimeRangeOptions),
	"ks_installed_app":         biliOpts(constants.KsInstalledAppSwitchOptions),
	"ks_region_type":           biliOpts(constants.KsRegionTypeOptions),
	"ks_region_user":           biliOpts(constants.KsRegionUserTypeOptions),
	"ks_share_user":            biliOpts(constants.KsShareUserOptions),
	"ks_behavior_type":         biliOpts(constants.KsBehaviorTypeOptions),
	"ks_behavior_scene":        biliOpts(constants.KsBehaviorSceneOptions),
	"ks_behavior_strength":     biliOpts(constants.KsBehaviorStrengthOptions),
	"ks_behavior_time":         biliOpts(constants.KsBehaviorTimeOptions),
	"ks_dmp_type":              biliOpts(constants.KsDmpTypeOptions),
	"ks_media_source":          biliOpts(constants.KsMediaSourceTypeOptions),
	"ks_intelli_type":          biliOpts(constants.KsIntelliTypeOptions),
	"ks_intelli_extend":        biliOpts(constants.KsIntelliExtendOptions),
	"ks_target_type":           biliOpts(constants.KsTargetTypeOptions),
	"ks_scene_inventory":       biliOpts(constants.KsSceneInventoryOptions),
	"ks_target_action":         biliOpts(constants.KsTargetActionOptions),
}

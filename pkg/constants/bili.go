package constants

// B站广告常量字典
// 数据来源：beiyu-backend/bm2 app/Constants/BilibiliConstants.php
// 说明：定向包维度（年龄/性别/网络/手机价位等）在原项目中来自 MongoDB bilibili_audience，
// 此处提供常用静态选项，接入真实数据后可在此调整。

type BiliOption struct {
	Value int
	Label string
}

type BiliStringOption struct {
	Value string
	Label string
}

// 推广目标
const (
	BiliPromotionPurposeSale = 2
	BiliPromotionPurposeApp  = 4
)

var BiliPromotionPurposeOptions = []BiliOption{
	{Value: BiliPromotionPurposeSale, Label: "销售线索收集"},
	{Value: BiliPromotionPurposeApp, Label: "应用推广"},
}

// 广告类型
const (
	BiliAdTypeAll = 0
)

var BiliAdTypeOptions = []BiliOption{
	{Value: BiliAdTypeAll, Label: "所有广告"},
}

// 投放类型
const (
	BiliDeliveryManual = 0
	BiliDeliveryAuto   = 1
)

var BiliDeliveryOptions = []BiliOption{
	{Value: BiliDeliveryManual, Label: "手动投放"},
	{Value: BiliDeliveryAuto, Label: "自动投放"},
}

// 预算类型
const (
	BiliBudgetSpecificDaily = 1
	BiliBudgetUnlimited     = 2
	BiliBudgetTotal         = 3
)

var BiliBudgetOptions = []BiliOption{
	{Value: BiliBudgetUnlimited, Label: "不限预算"},
	{Value: BiliBudgetSpecificDaily, Label: "指定日预算"},
	{Value: BiliBudgetTotal, Label: "总预算"},
}

// 投放方式
const (
	BiliSpeedModeStandard    = 1
	BiliSpeedModeAccelerated = 2
)

var BiliSpeedModeOptions = []BiliOption{
	{Value: BiliSpeedModeStandard, Label: "匀速投放"},
	{Value: BiliSpeedModeAccelerated, Label: "加速投放"},
}

// 推广内容
const (
	BiliPromotionContentClue     = 2
	BiliPromotionContentMiniGame = 65
)

var BiliPromotionContentOptions = []BiliOption{
	{Value: BiliPromotionContentClue, Label: "线索"},
	{Value: BiliPromotionContentMiniGame, Label: "小程序/小游戏"},
}

// 小游戏类型
const (
	BiliMiniGame       = 5
	BiliMiniWechatGame = 2
)

var BiliMiniGameOptions = []BiliOption{
	{Value: BiliMiniGame, Label: "B站小游戏"},
	{Value: BiliMiniWechatGame, Label: "微信小游戏"},
}

// 投放日期类型
const (
	BiliUnitDateNow = 0
	BiliUnitDateEnd = 1
)

var BiliUnitDateTypeOptions = []BiliOption{
	{Value: BiliUnitDateNow, Label: "从开始时间长期投放"},
	{Value: BiliUnitDateEnd, Label: "设置开始和结束日期"},
}

// 投放时间段类型
const (
	BiliUnitTimeNo  = 0
	BiliUnitTimeCom = 1
)

var BiliUnitTimeTypeOptions = []BiliOption{
	{Value: BiliUnitTimeNo, Label: "不限"},
	{Value: BiliUnitTimeCom, Label: "指定时间段"},
}

// 竞价策略
const (
	BiliStrategyStab = 0
	BiliStrategyMax  = 1
)

var BiliStrategyOptions = []BiliOption{
	{Value: BiliStrategyStab, Label: "稳定成本投放"},
	{Value: BiliStrategyMax, Label: "最大转化投放"},
}

// 出价方式
const (
	BiliBidOcpm = 1
	BiliBidCpc  = 2
	BiliBidCpm  = 3
)

var BiliBidOptions = []BiliOption{
	{Value: BiliBidOcpm, Label: "OCPM"},
	{Value: BiliBidCpm, Label: "CPM"},
	{Value: BiliBidCpc, Label: "CPC"},
}

// 优化目标
const (
	BiliGoldFormSubmit        = 4
	BiliGoldOrderSubmit       = 5
	BiliGoldPlay              = 9
	BiliGoldFirstPay          = 11
	BiliGoldEffectClues       = 13
	BiliGoldFormPay           = 18
	BiliGoldLinkClick         = 30
	BiliGoldDayPayRoi         = 32
	BiliGoldEveryPay          = 33
	BiliGoldWechatCopy        = 34
	BiliGoldFinish            = 35
	BiliGoldKeyBehavior       = 38
	BiliGoldComponentClick    = 50
	BiliGoldDayCash           = 54
	BiliGoldDayMixCash        = 57
	BiliGoldCredit            = 36
	BiliGoldCustomKeyBehavior = 60
	BiliGold7PayRoi           = 56
)

var BiliOptimizeGoldOptions = []BiliOption{
	{Value: BiliGoldFormSubmit, Label: "表单提交"},
	{Value: BiliGoldOrderSubmit, Label: "订单提交"},
	{Value: BiliGoldPlay, Label: "稿件播放"},
	{Value: BiliGoldFirstPay, Label: "应用内付费(首次)"},
	{Value: BiliGoldEffectClues, Label: "有效线索"},
	{Value: BiliGoldFormPay, Label: "表单付费"},
	{Value: BiliGoldLinkClick, Label: "评论链接点击"},
	{Value: BiliGoldDayPayRoi, Label: "24小时付费ROI"},
	{Value: BiliGoldEveryPay, Label: "每次付费"},
	{Value: BiliGoldWechatCopy, Label: "微信复制"},
	{Value: BiliGoldFinish, Label: "完件"},
	{Value: BiliGoldKeyBehavior, Label: "关键行为"},
	{Value: BiliGoldComponentClick, Label: "组件点击"},
	{Value: BiliGoldDayCash, Label: "24小时变现ROI"},
	{Value: BiliGoldDayMixCash, Label: "24小时混合变现ROI"},
	{Value: BiliGoldCredit, Label: "授信"},
	{Value: BiliGoldCustomKeyBehavior, Label: "自定义关键行为"},
	{Value: BiliGold7PayRoi, Label: "7日付费ROI"},
}

// 深度优化目标
var BiliDeepOptimizeGoldOptions = []BiliOption{
	{Value: BiliGoldDayPayRoi, Label: "24小时付费ROI"},
	{Value: BiliGoldEveryPay, Label: "每次付费"},
	{Value: BiliGoldFormPay, Label: "表单付费"},
	{Value: BiliGoldFirstPay, Label: "应用内付费(首次)"},
	{Value: BiliGoldCredit, Label: "授信"},
	{Value: BiliGold7PayRoi, Label: "7日付费ROI"},
}

// 深度优化方式
const (
	BiliDeepGoldCustom = 0
)

var BiliDeepGoldOptions = []BiliOption{
	{Value: BiliDeepGoldCustom, Label: "自定义双出价"},
}

// 投放模式
const (
	BiliLaunchNativeMerge = 3
)

var BiliLaunchOptions = []BiliOption{
	{Value: BiliLaunchNativeMerge, Label: "原生合并创意"},
}

// 流量类型
const (
	BiliInterAll    = 10040000
	BiliInterMobile = 10020000
	BiliInterPC     = 10010000
)

var BiliNetworkOptions = []BiliOption{
	{Value: BiliInterAll, Label: "全部"},
	{Value: BiliInterMobile, Label: "移动"},
	{Value: BiliInterPC, Label: "PC"},
}

// 自定义出价
const (
	BiliCustomBidNormal = "CUSTOM_BID_TYPE_NORMAL"
	BiliCustomBidRand   = "CUSTOM_BID_TYPE_RAND"
)

var BiliCustomBidOptions = []BiliStringOption{
	{Value: BiliCustomBidNormal, Label: "固定出价"},
	{Value: BiliCustomBidRand, Label: "随机出价"},
}

// 品牌类型
const (
	BiliBrandCustom = 2
)

var BiliBrandOptions = []BiliOption{
	{Value: BiliBrandCustom, Label: "自定义"},
}

// 地域类型
const (
	BiliAreaAll    = 0
	BiliAreaNow    = 1
	BiliAreaLive   = 2
	BiliAreaTravel = 3
)

var BiliAreaTypeOptions = []BiliOption{
	{Value: BiliAreaAll, Label: "全部"},
	{Value: BiliAreaNow, Label: "实时在此用户"},
	{Value: BiliAreaLive, Label: "常住在此用户"},
	{Value: BiliAreaTravel, Label: "旅游在此用户"},
}

// 安装过滤
const (
	BiliInstalledAll    = 0
	BiliInstalledFinish = 1
	BiliInstalledOther  = 2
	BiliInstalledNo     = 3
)

var BiliInstalledOptions = []BiliOption{
	{Value: BiliInstalledAll, Label: "不限"},
	{Value: BiliInstalledFinish, Label: "过滤已安装"},
	{Value: BiliInstalledOther, Label: "安装其他"},
	{Value: BiliInstalledNo, Label: "定向已安装"},
}

// 已安装应用
var BiliInstalledAppOptions = []BiliOption{
	{Value: 1, Label: "京东"},
	{Value: 2, Label: "淘宝"},
	{Value: 3, Label: "拼多多"},
	{Value: 4, Label: "微信"},
}

// 标签模糊
const (
	BiliTagFuzzyNo   = 0
	BiliTagFuzzySure = 1
	BiliTagFuzzyLike = 2
)

var BiliTagFuzzyOptions = []BiliOption{
	{Value: BiliTagFuzzyNo, Label: "无"},
	{Value: BiliTagFuzzySure, Label: "精准"},
	{Value: BiliTagFuzzyLike, Label: "模糊"},
}

// 智能词包类型
const (
	BiliTitleWordArea   = 1
	BiliTitleWordGender = 2
	BiliTitleWordDevice = 3
	BiliTitleWordAge    = 4
)

var BiliTitleWordOptions = []BiliOption{
	{Value: BiliTitleWordArea, Label: "智能地域"},
	{Value: BiliTitleWordGender, Label: "智能性别"},
	{Value: BiliTitleWordDevice, Label: "智能设备"},
	{Value: BiliTitleWordAge, Label: "智能年龄"},
}

var BiliTitleWordLimit = map[int]int{
	BiliTitleWordArea:   4,
	BiliTitleWordGender: 1,
	BiliTitleWordDevice: 2,
	BiliTitleWordAge:    2,
}

// 定向包常用维度（原项目来自 MongoDB bilibili_audience，此处为静态常用选项）
var BiliGenderOptions = []BiliOption{
	{Value: 1, Label: "男"},
	{Value: 2, Label: "女"},
}

var BiliAgeOptions = []BiliOption{
	{Value: 1, Label: "0-17岁"},
	{Value: 2, Label: "18-23岁"},
	{Value: 3, Label: "24-30岁"},
	{Value: 4, Label: "31-40岁"},
	{Value: 5, Label: "41-50岁"},
	{Value: 6, Label: "50岁以上"},
}

var BiliNetOptions = []BiliOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "WiFi"},
	{Value: 2, Label: "4G"},
	{Value: 3, Label: "3G"},
	{Value: 4, Label: "2G"},
}

var BiliPhonePriceOptions = []BiliOption{
	{Value: 1, Label: "0-999"},
	{Value: 2, Label: "1000-1999"},
	{Value: 3, Label: "2000-2999"},
	{Value: 4, Label: "3000-3999"},
	{Value: 5, Label: "4000-4999"},
	{Value: 6, Label: "5000以上"},
}

var BiliOsOptions = []BiliOption{
	{Value: 398, Label: "iOS"},
	{Value: 399, Label: "Android"},
	{Value: 421, Label: "iOS(iPad)"},
}

// 转化用户过滤
const (
	BiliConvertedNo   = 0
	BiliConvertedHide = 1
	BiliConvertedShow = 2
)

var BiliConvertedUserFilterOptions = []BiliOption{
	{Value: BiliConvertedNo, Label: "不限"},
	{Value: BiliConvertedHide, Label: "过滤已转化"},
	{Value: BiliConvertedShow, Label: "定向已转化"},
}

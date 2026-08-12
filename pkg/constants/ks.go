package constants

// 快手广告常量字典
// 数据来源：beiyu-backend/bm2 app/Constants/KuaiShouConstants.php
// 说明：枚举数组中的键为 self::XXX 命名常量，此处已解析为真实整数 id；
// 原枚举中 "不限" 使用 "ALL" 字符串键，统一映射为 0。

type KsOption = BiliOption
type KsStringOption = BiliStringOption

// 投放目标
const KsMarketTargetMiniProgram = 4

var KsMarketTargetOptions = []KsOption{
	{Value: KsMarketTargetMiniProgram, Label: "小程序推广"},
}

// 广告场景
const (
	KsSceneAppInstall        = 2
	KsSceneAppActive         = 7
	KsSceneFansLive          = 16
	KsSceneKwaiMiniProgram   = 19
	KsSceneWechatMiniProgram = 32
)

var KsSceneOptions = []KsOption{
	{Value: KsSceneAppInstall, Label: "应用下载"},
	{Value: KsSceneAppActive, Label: "应用唤起"},
	{Value: KsSceneKwaiMiniProgram, Label: "快手小程序/小游戏"},
	{Value: KsSceneWechatMiniProgram, Label: "微信小程序/小游戏"},
	{Value: KsSceneFansLive, Label: "粉丝/直播推广"},
}

// 广告类型
const (
	KsAdTypeShow    = 0
	KsAdTypeSearch  = 1
)

var KsAdTypeOptions = []KsOption{
	{Value: KsAdTypeShow, Label: "展示广告"},
	{Value: KsAdTypeSearch, Label: "搜索广告"},
}

// 投放方式（自动/常规）
const (
	KsAutoManagerNormal = 0
	KsAutoManagerAuto   = 1
)

var KsAutoManagerOptions = []KsOption{
	{Value: KsAutoManagerNormal, Label: "常规投放"},
	{Value: KsAutoManagerAuto, Label: "全自动投放"},
}

// 预算类型
const (
	KsBudgetNoLimit   = 1
	KsBudgetDay       = 2
	KsBudgetEveryDay  = 3
)

var KsBudgetTypeOptions = []KsOption{
	{Value: KsBudgetNoLimit, Label: "不限"},
	{Value: KsBudgetDay, Label: "统一预算"},
	{Value: KsBudgetEveryDay, Label: "分日预算"},
}

// 竞价策略
const (
	KsBidStrategyStab = 0
	KsBidStrategyMax  = 1
)

var KsBidStrategyOptions = []KsOption{
	{Value: KsBidStrategyStab, Label: "稳定成本"},
	{Value: KsBidStrategyMax, Label: "最大转化"},
}

// 素材优选
const (
	KsAutoPhotoScopeSys = 0
)

var KsAutoPhotoScopeOptions = []KsOption{
	{Value: KsAutoPhotoScopeSys, Label: "系统优选"},
}

// 创意类型
const (
	KsCreativeUnitAdvanced = 7
	KsCreativeUnitCustom   = 4
)

var KsCreativeUnitOptions = []KsOption{
	{Value: KsCreativeUnitAdvanced, Label: "程序化创意"},
	{Value: KsCreativeUnitCustom, Label: "自定义创意"},
}

// 小游戏类型
const (
	KsMiniKwaiGame      = 2
	KsMiniKwaiProgram   = 1
	KsMiniWechatGame    = 4
	KsMiniWechatProgram = 3
)

var KsMiniTypeOptions = []KsOption{
	{Value: KsMiniKwaiGame, Label: "快手小游戏"},
	{Value: KsMiniKwaiProgram, Label: "快手小程序"},
	{Value: KsMiniWechatGame, Label: "微信小游戏"},
	{Value: KsMiniWechatProgram, Label: "微信小程序"},
}

// 投放日期类型
const (
	KsUnitDateNow = 0
	KsUnitDateEnd = 1
)

var KsUnitDateTypeOptions = []KsOption{
	{Value: KsUnitDateNow, Label: "从开始时间长期投放"},
	{Value: KsUnitDateEnd, Label: "设置开始和结束日期"},
}

// 投放时间段类型
const (
	KsUnitTimeNo  = 0
	KsUnitTimeCom = 1
)

var KsUnitTimeTypeOptions = []KsOption{
	{Value: KsUnitTimeNo, Label: "不限"},
	{Value: KsUnitTimeCom, Label: "指定时间段"},
}

// 出价方式
const (
	KsBidOcpm = 10
	KsBidCpm  = 1
	KsBidCpc  = 2
	KsBidMcb  = 12
)

var KsBidTypeOptions = []KsOption{
	{Value: KsBidOcpm, Label: "OCPM"},
	{Value: KsBidCpm, Label: "CPM"},
	{Value: KsBidCpc, Label: "CPC"},
	{Value: KsBidMcb, Label: "MCB"},
}

// 优化目标
const (
	KsGoldCoverClick = 1
	KsGoldBehavior   = 2
	KsGoldActivate   = 180
	KsGoldPay        = 190
	KsGoldRoi        = 191
)

var KsOcpxActionOptions = []KsOption{
	{Value: KsGoldCoverClick, Label: "封面点击数"},
	{Value: KsGoldBehavior, Label: "行为数"},
	{Value: KsGoldActivate, Label: "激活"},
	{Value: KsGoldPay, Label: "付费"},
	{Value: KsGoldRoi, Label: "首日ROI"},
}

// 深度优化目标
const (
	KsDeepNone        = 0
	KsDeepPay         = 3
	KsDeepRetention7  = 7
	KsDeepFinish      = 10
	KsDeepCredit      = 11
	KsDeepCart        = 13
	KsDeepOrder       = 14
	KsDeepBuy         = 15
	KsDeepRoi         = 92
)

var KsDeepConversionOptions = []KsOption{
	{Value: KsDeepNone, Label: "无"},
	{Value: KsDeepPay, Label: "付费"},
	{Value: KsDeepRetention7, Label: "次日留存"},
	{Value: KsDeepFinish, Label: "完件"},
	{Value: KsDeepCredit, Label: "授信"},
	{Value: KsDeepCart, Label: "添加购物车"},
	{Value: KsDeepOrder, Label: "提交订单"},
	{Value: KsDeepBuy, Label: "购买"},
	{Value: KsDeepRoi, Label: "首日ROI"},
}

// 展现方式
const (
	KsShowRandom = 1
	KsShowSmart  = 2
)

var KsShowModeOptions = []KsOption{
	{Value: KsShowRandom, Label: "随机轮播"},
	{Value: KsShowSmart, Label: "智能优选"},
}

// 自定义出价方式
const (
	KsBidWayDual   = 0
	KsBidWayUnit   = 1
)

var KsBidWayOptions = []KsOption{
	{Value: KsBidWayDual, Label: "双出价模式"},
	{Value: KsBidWayUnit, Label: "付费单价（白名单可用）"},
}

// 投放速度
const (
	KsSpeedNormal = 1
	KsSpeedSmooth = 2
	KsSpeedLow    = 3
)

var KsSpeedTypeOptions = []KsOption{
	{Value: KsSpeedNormal, Label: "正常投放"},
	{Value: KsSpeedSmooth, Label: "平滑投放"},
	{Value: KsSpeedLow, Label: "优先低成本"},
}

// 手动/自动出价
const (
	KsSmartBidManual = 0
	KsSmartBidAuto   = 1
)

var KsSmartBidOptions = []KsOption{
	{Value: KsSmartBidManual, Label: "手动出价"},
	{Value: KsSmartBidAuto, Label: "自动出价"},
}

// ---------- 定向包 ----------

// 性别
var KsGenderOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "女性"},
	{Value: 2, Label: "男性"},
}

// 年龄
var KsAgeOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "选择年龄段"},
	{Value: 2, Label: "自定义"},
}

// 网络环境
var KsNetworkOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "Wi-Fi"},
	{Value: 2, Label: "移动网络"},
}

// 运营商
var KsOperatorsOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "中国移动"},
	{Value: 2, Label: "中国电信"},
	{Value: 3, Label: "中国联通"},
}

// 操作系统
var KsPlatformOsOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "Android"},
	{Value: 2, Label: "IOS"},
	{Value: 3, Label: "Android&iOS"},
	{Value: 4, Label: "鸿蒙"},
}

// 设备价位类型
const (
	KsDevicePriceNo     = 0
	KsDevicePriceCustom = 1
)

var KsDevicePriceTypeOptions = []KsOption{
	{Value: KsDevicePriceNo, Label: "不限"},
	{Value: KsDevicePriceCustom, Label: "自定义"},
}

// 设备价位
var KsDevicePriceOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 11, Label: "1000元以下"},
	{Value: 12, Label: "1001~1500"},
	{Value: 1, Label: "1500元以下"},
	{Value: 2, Label: "1501～2000"},
	{Value: 3, Label: "2001～2500"},
	{Value: 4, Label: "2501～3000"},
	{Value: 5, Label: "3001～3500"},
	{Value: 6, Label: "3501～4000"},
	{Value: 7, Label: "4001～4500"},
	{Value: 8, Label: "4501～5000"},
	{Value: 9, Label: "5001～5500"},
	{Value: 10, Label: "5500元以上"},
	{Value: 13, Label: "4001~5000"},
	{Value: 14, Label: "5001~6000"},
	{Value: 15, Label: "6000元以上"},
}

// APP兴趣方式
var KsAppInterestTypeOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "按分类"},
	{Value: 2, Label: "按名称"},
}

// 过滤已转化层级
var KsFilterConvertedLevelOptions = []KsOption{
	{Value: 0, Label: "不限(默认)"},
	{Value: 1, Label: "广告组"},
	{Value: 2, Label: "广告计划"},
	{Value: 3, Label: "本账户"},
	{Value: 4, Label: "公司主体"},
	{Value: 5, Label: "APP"},
}

// 转化时间
var KsFilterTimeRangeOptions = []KsOption{
	{Value: 0, Label: "30天"},
	{Value: 1, Label: "60天"},
	{Value: 2, Label: "90天"},
}

// 排除已安装
var KsInstalledAppSwitchOptions = []KsOption{
	{Value: 0, Label: "过滤"},
	{Value: 1, Label: "不限"},
}

// 地域类型
var KsRegionTypeOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "指定地区"},
}

// 地域用户
var KsRegionUserTypeOptions = []KsOption{
	{Value: 0, Label: "实时地用户"},
	{Value: 1, Label: "常驻地用户"},
	{Value: 2, Label: "该地区所有用户(常驻地+居住地+实时地)"},
	{Value: 4, Label: "到此旅游用户"},
	{Value: 5, Label: "居住地用户"},
}

// 共享用户过滤
var KsShareUserOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "过滤"},
}

// 行为定向
var KsBehaviorTypeOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "自定义"},
}

// 行为场景
var KsBehaviorSceneOptions = []KsOption{
	{Value: 1, Label: "社区"},
	{Value: 2, Label: "APP"},
	{Value: 3, Label: "电商"},
	{Value: 4, Label: "推广"},
}

// 行为强度
var KsBehaviorStrengthOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "高强度"},
}

// 行为时间
var KsBehaviorTimeOptions = []KsOption{
	{Value: 0, Label: "7天"},
	{Value: 1, Label: "15天"},
	{Value: 2, Label: "30天"},
	{Value: 3, Label: "90天"},
	{Value: 4, Label: "180天"},
}

// 人群包类型
var KsDmpTypeOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "自定义"},
	{Value: 2, Label: "种子人群"},
}

// 媒体来源
var KsMediaSourceTypeOptions = []KsOption{
	{Value: 0, Label: "不限"},
	{Value: 1, Label: "行业优质流量包"},
}

// 智能放量开关
const (
	KsIntelliOff = 0
	KsIntelliOn  = 1
)

var KsIntelliTypeOptions = []KsOption{
	{Value: KsIntelliOff, Label: "关闭"},
	{Value: KsIntelliOn, Label: "开启"},
}

// 智能放量保护
var KsIntelliExtendOptions = []KsOption{
	{Value: 1, Label: "不可突破年龄"},
	{Value: 2, Label: "不可突破性别"},
	{Value: 3, Label: "不可突破地域"},
}

// 定向类型
const (
	KsTargetTypeDir  = 0
	KsTargetTypeExcl = 1
	KsTargetTypeBoth = 2
)

var KsTargetTypeOptions = []KsOption{
	{Value: KsTargetTypeDir, Label: "定向"},
	{Value: KsTargetTypeExcl, Label: "排除"},
	{Value: KsTargetTypeBoth, Label: "同时定向排除"},
}

// 投放位置（scene_id）
const (
	KsSceneInventoryPreferred = 1
	KsSceneInventorySlideFeed = 2
	KsSceneInventoryVideoPage = 3
	KsSceneInventoryUnion     = 5
	KsSceneInventoryUpDown    = 6
	KsSceneInventoryFeed      = 7
	KsSceneInventoryUpDownMix = 8
)

var KsSceneInventoryOptions = []KsOption{
	{Value: KsSceneInventoryPreferred, Label: "快手-优选广告位"},
	{Value: KsSceneInventorySlideFeed, Label: "快手-下滑大屏广告+信息流广告"},
	{Value: KsSceneInventoryVideoPage, Label: "快手-视频播放页广告/便利贴广告"},
	{Value: KsSceneInventoryUnion, Label: "快手-联盟广告"},
	{Value: KsSceneInventoryUpDown, Label: "快手-上下滑大屏广告"},
	{Value: KsSceneInventoryFeed, Label: "快手-信息流广告"},
	{Value: KsSceneInventoryUpDownMix, Label: "快手-上下滑大屏广告+信息流广告"},
}

// 目标行为（cpa_target 动作）
const (
	KsTargetActionIos = 2
	KsTargetActionAndroidDownload  = 30
	KsTargetActionAndroidFinish    = 31
	KsTargetActionAndroidInstall   = 32
)

var KsTargetActionOptions = []KsOption{
	{Value: KsTargetActionIos, Label: "检测到用户行为（actionBar 点击）-iOS"},
	{Value: KsTargetActionAndroidDownload, Label: "开始下载后-安卓"},
	{Value: KsTargetActionAndroidFinish, Label: "下载完成后-安卓"},
	{Value: KsTargetActionAndroidInstall, Label: "安装完成后-安卓"},
}

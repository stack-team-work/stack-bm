package constants

// 头条广告常量字典
// 数据来源：beiyu-backend/bm2 app/Constants/ByteDanceV2Constants.php
// 说明：头条模板枚举大多为字符串键（如 'APP'、'AD_CONVERT_TYPE_PAY'），
// 与 b站/快手 使用整数 id 不同，故提供字符串型选项。

type TtOption = BiliOption
type TtStringOption = BiliStringOption

// ========== 广告模板（项目模板） ==========

// 推广目的
const (
	TtLandingApp          = "APP"
	TtLandingMicroGame    = "MICRO_GAME"
	TtLandingLink         = "LINK"
	TtLandingNativeAction = "NATIVE_ACTION"
)

var TtLandingTypeOptions = []TtStringOption{
	{Value: TtLandingApp, Label: "应用推广"},
	{Value: TtLandingMicroGame, Label: "小程序"},
	{Value: TtLandingLink, Label: "销售线索推广"},
	{Value: TtLandingNativeAction, Label: "原生互动"},
}

// 子目标
const (
	TtAppPromotionDownload = "DOWNLOAD"
	TtAppPromotionLaunch   = "LAUNCH"
	TtAppPromotionReserve  = "RESERVE"
)

var TtAppPromotionTypeOptions = []TtStringOption{
	{Value: TtAppPromotionDownload, Label: "应用下载"},
	{Value: TtAppPromotionLaunch, Label: "应用调用"},
	{Value: TtAppPromotionReserve, Label: "预约下载"},
}

// 小程序类型
const (
	TtMicroPromotionWechatGame = "WECHAT_GAME"
	TtMicroPromotionWechatApp  = "WECHAT_APP"
	TtMicroPromotionByteGame   = "BYTE_GAME"
	TtMicroPromotionByteApp    = "BYTE_APP"
)

var TtMicroPromotionTypeOptions = []TtStringOption{
	{Value: TtMicroPromotionWechatGame, Label: "微信小游戏"},
	{Value: TtMicroPromotionWechatApp, Label: "微信小程序"},
	{Value: TtMicroPromotionByteGame, Label: "抖音小游戏"},
	{Value: TtMicroPromotionByteApp, Label: "抖音小程序"},
}

// 投放模式
const (
	TtDeliveryManual     = "MANUAL"
	TtDeliveryProcedural = "PROCEDURAL"
)

var TtDeliveryModeOptions = []TtStringOption{
	{Value: TtDeliveryManual, Label: "手动投放"},
	{Value: TtDeliveryProcedural, Label: "自动投放"},
}

// 营销场景
const (
	TtMarketingGoalVideoAndImage = "VIDEO_AND_IMAGE"
	TtMarketingGoalLive          = "LIVE"
)

var TtMarketingGoalOptions = []TtStringOption{
	{Value: TtMarketingGoalVideoAndImage, Label: "短视频/图片"},
	{Value: TtMarketingGoalLive, Label: "直播"},
}

// 广告类型
const (
	TtAdTypeAll    = "ALL"
	TtAdTypeSearch = "SEARCH"
)

var TtAdTypeOptions = []TtStringOption{
	{Value: TtAdTypeAll, Label: "通投广告"},
	{Value: TtAdTypeSearch, Label: "搜索广告"},
}

// 下载方式
const (
	TtDownloadTypeDownloadUrl = "DOWNLOAD_URL"
	TtDownloadTypeExternalUrl = "EXTERNAL_URL"
)

var TtDownloadTypeOptions = []TtStringOption{
	{Value: TtDownloadTypeDownloadUrl, Label: "直接下载"},
	{Value: TtDownloadTypeExternalUrl, Label: "落地页下载"},
}

// 下载模式
const (
	TtDownloadModeDefault          = "DEFAULT"
	TtDownloadModeAppStoreDelivery = "APP_STORE_DELIVERY"
)

var TtDownloadModeOptions = []TtStringOption{
	{Value: TtDownloadModeDefault, Label: "默认下载"},
	{Value: TtDownloadModeAppStoreDelivery, Label: "优先商店下载"},
}

// 竞价策略
const (
	TtBidTypeCustom       = "CUSTOM"
	TtBidTypeUpperControl = "UPPER_CONTROL"
	TtBidTypeNoBid        = "NO_BID"
)

var TtBidTypeOptions = []TtStringOption{
	{Value: TtBidTypeCustom, Label: "稳定成本"},
	{Value: TtBidTypeUpperControl, Label: "最优成本"},
	{Value: TtBidTypeNoBid, Label: "最大转化投放"},
}

// 预算择优分配
const (
	TtBudgetOptimizeOff = "OFF"
	TtBudgetOptimizeOn  = "ON"
)

var TtBudgetOptimizeSwitchOptions = []TtStringOption{
	{Value: TtBudgetOptimizeOff, Label: "不开启"},
	{Value: TtBudgetOptimizeOn, Label: "开启"},
}

// 优化目标
const (
	TtAdConvertActive          = "AD_CONVERT_TYPE_ACTIVE"
	TtAdConvertActiveRegister  = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	TtAdConvertPay             = "AD_CONVERT_TYPE_PAY"
	TtAdConvertGameAddiction   = "AD_CONVERT_TYPE_GAME_ADDICTION"
	TtAdConvertPurchaseRoi     = "AD_CONVERT_TYPE_PURCHASE_ROI"
	TtAdConvertNotifyDownload  = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	TtAdConvertClickNum        = "AD_CONVERT_TYPE_CLICK_NUM"
	TtAdConvertShowOffNum      = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	TtAdConvertRss             = "AD_CONVERT_TYPE_RSS"
	TtAdConvertLiveGiftAction  = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	TtAdConvertLiveEnterAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	TtAdConvertLiveStayTime    = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	TtAdConvertLiveComClick    = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	TtAdConvertLiveJoinGroup   = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
)

var TtAdConvertTypeOptions = []TtStringOption{
	{Value: TtAdConvertActive, Label: "激活"},
	{Value: TtAdConvertActiveRegister, Label: "注册"},
	{Value: TtAdConvertPay, Label: "付费"},
	{Value: TtAdConvertGameAddiction, Label: "关键行为"},
	{Value: TtAdConvertNotifyDownload, Label: "预约下载"},
	{Value: TtAdConvertClickNum, Label: "点击量"},
	{Value: TtAdConvertShowOffNum, Label: "展示量"},
	{Value: TtAdConvertRss, Label: "账号关注"},
	{Value: TtAdConvertLiveGiftAction, Label: "直播间内打赏"},
	{Value: TtAdConvertLiveEnterAction, Label: "直播间观看"},
	{Value: TtAdConvertLiveStayTime, Label: "直播间停留"},
	{Value: TtAdConvertLiveComClick, Label: "组件点击"},
	{Value: TtAdConvertLiveJoinGroup, Label: "粉丝入群"},
}

// 深度优化目标
const (
	TtDeepConvertNone          = "NONE"
	TtDeepConvertPurchaseRoi   = "AD_CONVERT_TYPE_PURCHASE_ROI"
	TtDeepConvertUgRoi         = "AD_CONVERT_TYPE_UG_ROI"
	TtDeepConvertNextDayOpen   = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	TtDeepConvertRetention7d   = "AD_CONVERT_TYPE_RETENTION_7D"
	TtDeepConvertLtRoi         = "AD_CONVERT_TYPE_LT_ROI"
	TtDeepConvertPurchaseRoi7d = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	TtDeepConvertLtvRoiSeven   = "AD_CONVERT_TYPE_LTV_ROI_SEVEN"
)

var TtDeepAdConvertTypeOptions = []TtStringOption{
	{Value: TtDeepConvertNone, Label: "无"},
	{Value: TtDeepConvertPurchaseRoi, Label: "付费ROI"},
	{Value: TtDeepConvertUgRoi, Label: "ROI三目标"},
	{Value: TtDeepConvertNextDayOpen, Label: "次留"},
	{Value: TtDeepConvertRetention7d, Label: "7日留存"},
	{Value: TtDeepConvertLtRoi, Label: "广告变现"},
	{Value: TtDeepConvertPurchaseRoi7d, Label: "付费roi-7日"},
	{Value: TtDeepConvertLtvRoiSeven, Label: "广告变现-7日"},
}

// 深度出价方式
const (
	TtDeepBidNone           = "NONE"
	TtDeepBidMin            = "DEEP_BID_MIN"
	TtDeepBidRoiCoefficient = "ROI_COEFFICIENT"
	TtDeepBidRoiPacing      = "ROI_PACING"
	TtDeepBidPerAction      = "BID_PER_ACTION"
	TtDeepBidSocialRoi      = "SOCIAL_ROI"
	TtDeepBidPerAndSevenRoi = "PER_AND_SEVEN_PAY_ROI"
	TtDeepBidSevenRoiCoeff  = "SEVEN_ROI_COEFFICIENT"
)

var TtDeepBidTypeOptions = []TtStringOption{
	{Value: TtDeepBidNone, Label: "首次付费出价"},
	{Value: TtDeepBidMin, Label: "自定义双出价"},
	{Value: TtDeepBidRoiCoefficient, Label: "ROI系数出价"},
	{Value: TtDeepBidRoiPacing, Label: "ROI系数(自动优化)"},
	{Value: TtDeepBidPerAction, Label: "每次付费出价"},
	{Value: TtDeepBidSocialRoi, Label: "ROI三出价"},
	{Value: TtDeepBidPerAndSevenRoi, Label: "每次付费+7日ROI"},
	{Value: TtDeepBidSevenRoiCoeff, Label: "7日ROI"},
}

// 广告位置大类
const (
	TtInventoryCatalogManual         = "MANUAL"
	TtInventoryCatalogUniversalSmart = "UNIVERSAL_SMART"
)

var TtInventoryCatalogOptions = []TtStringOption{
	{Value: TtInventoryCatalogManual, Label: "首选媒体"},
	{Value: TtInventoryCatalogUniversalSmart, Label: "通投智选"},
}

// 投放位置
const (
	TtInventoryFeed           = "INVENTORY_FEED"
	TtInventoryVideoFeed      = "INVENTORY_VIDEO_FEED"
	TtInventoryHotsoonFeed    = "INVENTORY_HOTSOON_FEED"
	TtInventoryAwemeFeed      = "INVENTORY_AWEME_FEED"
	TtInventoryTomatoNovel    = "INVENTORY_TOMATO_NOVEL"
	TtInventoryUnionSlot      = "INVENTORY_UNION_SLOT"
	TtInventoryUnionBoutique  = "UNION_BOUTIQUE_GAME"
	TtInventorySearch         = "INVENTORY_SEARCH"
	TtInventoryHomedAggregate = "INVENTORY_HOMED_AGGREGATE"
)

var TtInventoryTypeOptions = []TtStringOption{
	{Value: TtInventoryFeed, Label: "今日头条"},
	{Value: TtInventoryVideoFeed, Label: "西瓜视频"},
	{Value: TtInventoryHotsoonFeed, Label: "抖音火山版"},
	{Value: TtInventoryAwemeFeed, Label: "抖音短视频"},
	{Value: TtInventoryTomatoNovel, Label: "番茄小说"},
	{Value: TtInventoryUnionSlot, Label: "穿山甲"},
	{Value: TtInventoryUnionBoutique, Label: "ohayoo精品游戏"},
	{Value: TtInventorySearch, Label: "搜索广告"},
	{Value: TtInventoryHomedAggregate, Label: "住小帮"},
}

// 投放形式
const (
	TtUnionVideoOriginal = "ORIGINAL_VIDEO"
	TtUnionVideoRewarded = "REWARDED_VIDEO"
	TtUnionVideoSplash   = "SPLASH_VIDEO"
)

var TtUnionVideoTypeOptions = []TtStringOption{
	{Value: TtUnionVideoOriginal, Label: "原生视频"},
	{Value: TtUnionVideoRewarded, Label: "激励视频"},
	{Value: TtUnionVideoSplash, Label: "开屏视频"},
}

// 素材类型
const (
	TtMaterialsLive             = "LIVE_MATERIALS"
	TtMaterialsPromotion        = "PROMOTION_MATERIALS"
	TtMaterialsLiveMixPromotion = "LIVE_MIX_PROMOTION_MATERIALS"
)

var TtMaterialsTypeOptions = []TtStringOption{
	{Value: TtMaterialsLive, Label: "直播素材"},
	{Value: TtMaterialsPromotion, Label: "广告素材"},
	{Value: TtMaterialsLiveMixPromotion, Label: "直播素材与广告素材混投"},
}

// 投放时间类型
const (
	TtScheduleFromNow  = "SCHEDULE_FROM_NOW"
	TtScheduleStartEnd = "SCHEDULE_START_END"
)

var TtScheduleTypeOptions = []TtStringOption{
	{Value: TtScheduleFromNow, Label: "从今天起长期投放"},
	{Value: TtScheduleStartEnd, Label: "设置开始和结束日期"},
}

// 投放时段类型
const (
	TtScheduleTimeAllDay   = 0
	TtScheduleTimeSpecific = 1
)

var TtScheduleTimeTypeOptions = []TtOption{
	{Value: TtScheduleTimeAllDay, Label: "不限"},
	{Value: TtScheduleTimeSpecific, Label: "指定时间段"},
}

// 项目预算类型
const (
	TtBudgetModeInfinite = "BUDGET_MODE_INFINITE"
	TtBudgetModeDay      = "BUDGET_MODE_DAY"
)

var TtBudgetModeOptions = []TtStringOption{
	{Value: TtBudgetModeInfinite, Label: "不限"},
	{Value: TtBudgetModeDay, Label: "指定预算"},
}

// 自定义出价方式
const (
	TtCustomBidNormal = "CUSTOM_BID_TYPE_NORMAL"
	TtCustomBidLadder = "CUSTOM_BID_TYPE_LADDER"
	TtCustomBidRand   = "CUSTOM_BID_TYPE_RAND"
)

var TtCustomBidTypeOptions = []TtStringOption{
	{Value: TtCustomBidNormal, Label: "固定出价"},
	{Value: TtCustomBidLadder, Label: "阶梯出价"},
	{Value: TtCustomBidRand, Label: "随机出价"},
}

// 原生锚点
const (
	TtAnchorAuto     = "AUTO"
	TtAnchorOff      = "OFF"
	TtAnchorSelected = "SELECT"
)

var TtAnchorRelatedTypeOptions = []TtStringOption{
	{Value: TtAnchorAuto, Label: "自动生成"},
	{Value: TtAnchorOff, Label: "不启用"},
	{Value: TtAnchorSelected, Label: "手动选择"},
}

// 关键词匹配类型
const (
	TtKeywordsMatchPhrase    = "PHRASE"
	TtKeywordsMatchExtensive = "EXTENSIVE"
	TtKeywordsMatchPrecision = "PRECISION"
)

var TtKeywordsMatchTypeOptions = []TtStringOption{
	{Value: TtKeywordsMatchPhrase, Label: "短语匹配"},
	{Value: TtKeywordsMatchExtensive, Label: "广泛匹配"},
	{Value: TtKeywordsMatchPrecision, Label: "精准匹配"},
}

// 智能拓流
const (
	TtAutoExtendTrafficOn  = "ON"
	TtAutoExtendTrafficOff = "OFF"
)

var TtAutoExtendTrafficOptions = []TtStringOption{
	{Value: TtAutoExtendTrafficOn, Label: "开启"},
	{Value: TtAutoExtendTrafficOff, Label: "关闭"},
}

// 投放内容
const (
	TtPromotionAwemeHomePage   = "AWEME_HOME_PAGE"
	TtPromotionLandingPageLink = "LANDING_PAGE_LINK"
)

var TtPromotionTypeOptions = []TtStringOption{
	{Value: TtPromotionAwemeHomePage, Label: "抖音主页"},
	{Value: TtPromotionLandingPageLink, Label: "落地页"},
}

// 投放类型（自定义）
const (
	TtMktAdTypeDefault = 1
	TtMktAdTypeTtStar  = 2
)

var TtMktAdTypeOptions = []TtOption{
	{Value: TtMktAdTypeDefault, Label: "常规"},
	{Value: TtMktAdTypeTtStar, Label: "星广联投"},
}

// 项目成本稳投
const (
	TtProjectCustomOn  = "ON"
	TtProjectCustomOff = "OFF"
)

var TtProjectCustomOptions = []TtStringOption{
	{Value: TtProjectCustomOn, Label: "开启"},
	{Value: TtProjectCustomOff, Label: "不开启"},
}

// ========== 定向包模板 ==========

// 定向包推广类型
const (
	TtAudienceLandingAndroid  = "APP_ANDROID"
	TtAudienceLandingIos      = "APP_IOS"
	TtAudienceLandingExternal = "EXTERNAL"
)

var TtAudienceLandingTypeOptions = []TtStringOption{
	{Value: TtAudienceLandingAndroid, Label: "应用下载-安卓"},
	{Value: TtAudienceLandingIos, Label: "应用下载-IOS"},
	{Value: TtAudienceLandingExternal, Label: "落地页"},
}

// 地域类型
const (
	TtDistrictCity    = "CITY"
	TtDistrictCounty  = "COUNTY"
	TtDistrictRegion  = "REGION"
	TtDistrictOversea = "OVERSEA"
	TtDistrictNone    = "NONE"
)

var TtDistrictTypeOptions = []TtStringOption{
	{Value: TtDistrictCity, Label: "省市"},
	{Value: TtDistrictCounty, Label: "区县"},
	{Value: TtDistrictRegion, Label: "行政区域"},
	{Value: TtDistrictOversea, Label: "海外区域"},
	{Value: TtDistrictNone, Label: "不限"},
}

// 位置类型
const (
	TtLocationCurrent = "CURRENT"
	TtLocationHome    = "HOME"
	TtLocationTravel  = "TRAVEL"
	TtLocationAll     = "ALL"
)

var TtLocationTypeOptions = []TtStringOption{
	{Value: TtLocationCurrent, Label: "正在该地区的用户"},
	{Value: TtLocationHome, Label: "居住在该地区的用户"},
	{Value: TtLocationTravel, Label: "到该地区旅行的用户"},
	{Value: TtLocationAll, Label: "该地区内的所有用户"},
}

// 性别
const (
	TtGenderNone   = "NONE"
	TtGenderMale   = "GENDER_MALE"
	TtGenderFemale = "GENDER_FEMALE"
)

var TtGenderOptions = []TtStringOption{
	{Value: TtGenderNone, Label: "不限"},
	{Value: TtGenderMale, Label: "男"},
	{Value: TtGenderFemale, Label: "女"},
}

// 年龄
const (
	TtAge18_23 = "AGE_BETWEEN_18_23"
	TtAge24_30 = "AGE_BETWEEN_24_30"
	TtAge31_40 = "AGE_BETWEEN_31_40"
	TtAge41_49 = "AGE_BETWEEN_41_49"
	TtAge50    = "AGE_ABOVE_50"
)

var TtAgeOptions = []TtStringOption{
	{Value: TtAge18_23, Label: "18-23岁"},
	{Value: TtAge24_30, Label: "24-30岁"},
	{Value: TtAge31_40, Label: "31-40岁"},
	{Value: TtAge41_49, Label: "41-49岁"},
	{Value: TtAge50, Label: "大于等于50岁"},
}

// 职业
const (
	TtCareerCollegeStudent = "COLLEGE_STUDENT"
	TtCareerTeacher        = "TEACHER"
	TtCareerIt             = "IT"
	TtCareerCivilServants  = "CIVIL_SERVANTS"
	TtCareerFinancial      = "FINANCIAL"
	TtCareerMedicalStaff   = "MEDICAL_STAFF"
)

var TtCareerOptions = []TtStringOption{
	{Value: TtCareerCollegeStudent, Label: "大学生"},
	{Value: TtCareerTeacher, Label: "教师"},
	{Value: TtCareerIt, Label: "IT"},
	{Value: TtCareerCivilServants, Label: "公务员"},
	{Value: TtCareerFinancial, Label: "金融"},
	{Value: TtCareerMedicalStaff, Label: "医务人员"},
}

// 行为兴趣模式
const (
	TtInterestActionUnlimited = "UNLIMITED"
	TtInterestActionCustom    = "CUSTOM"
	TtInterestActionRecommend = "RECOMMEND"
)

var TtInterestActionModeOptions = []TtStringOption{
	{Value: TtInterestActionUnlimited, Label: "不限"},
	{Value: TtInterestActionCustom, Label: "自定义"},
	{Value: TtInterestActionRecommend, Label: "系统推荐"},
}

// 行为场景
const (
	TtActionSceneCommerce = "E-COMMERCE"
	TtActionSceneNews     = "NEWS"
	TtActionSceneApp      = "APP"
	TtActionSceneSearch   = "SEARCH"
)

var TtActionSceneOptions = []TtStringOption{
	{Value: TtActionSceneCommerce, Label: "电商互动行为"},
	{Value: TtActionSceneNews, Label: "资讯互动行为"},
	{Value: TtActionSceneApp, Label: "APP推广互动行为"},
	{Value: TtActionSceneSearch, Label: "搜索互动行为"},
}

// 行为天数
var TtActionDaysOptions = []TtOption{
	{Value: 7, Label: "7天"},
	{Value: 15, Label: "15天"},
	{Value: 30, Label: "30天"},
	{Value: 60, Label: "60天"},
	{Value: 90, Label: "90天"},
	{Value: 180, Label: "180天"},
	{Value: 365, Label: "365天"},
}

// 抖音达人用户行为
const (
	TtAwemeFanFollowed        = "FOLLOWED_USER"
	TtAwemeFanCommented       = "COMMENTED_USER"
	TtAwemeFanLiked           = "LIKED_USER"
	TtAwemeFanShared          = "SHARED_USER"
	TtAwemeFanLiveWatch       = "LIVE_WATCH"
	TtAwemeFanLiveEffective   = "LIVE_EFFECTIVE_WATCH"
	TtAwemeFanLiveComment     = "LIVE_COMMENT"
	TtAwemeFanLiveExceptional = "LIVE_EXCEPTIONAL"
	TtAwemeFanLiveGoodsClick  = "LIVE_GOODS_CLICK"
	TtAwemeFanLiveGoodsOrder  = "LIVE_GOODS_ORDER"
	TtAwemeFanGoodsCartsClick = "GOODS_CARTS_CLICK"
	TtAwemeFanGoodsCartsOrder = "GOODS_CARTS_ORDER"
)

var TtAwemeFanBehaviorsOptions = []TtStringOption{
	{Value: TtAwemeFanFollowed, Label: "关注"},
	{Value: TtAwemeFanCommented, Label: "视频互动-已评论用户"},
	{Value: TtAwemeFanLiked, Label: "视频互动-已点赞用户"},
	{Value: TtAwemeFanShared, Label: "视频互动-已分享用户"},
	{Value: TtAwemeFanLiveWatch, Label: "直播互动-观看"},
	{Value: TtAwemeFanLiveEffective, Label: "直播互动-有效观看"},
	{Value: TtAwemeFanLiveComment, Label: "直播互动-直播评论"},
	{Value: TtAwemeFanLiveExceptional, Label: "直播互动-打赏"},
	{Value: TtAwemeFanLiveGoodsClick, Label: "直播互动-商品点击"},
	{Value: TtAwemeFanLiveGoodsOrder, Label: "直播互动-商品下单"},
	{Value: TtAwemeFanGoodsCartsClick, Label: "直播互动-购物车点击"},
	{Value: TtAwemeFanGoodsCartsOrder, Label: "直播互动-购物车下单"},
}

// 抖音达人互动行为时间范围
const (
	TtAwemeFanTime15Days = "FIFTEEN_DAYS"
	TtAwemeFanTime30Days = "THIRTY_DAYS"
	TtAwemeFanTime60Days = "SIXTY_DAYS"
)

var TtAwemeFanTimeScopeOptions = []TtStringOption{
	{Value: TtAwemeFanTime15Days, Label: "15天"},
	{Value: TtAwemeFanTime30Days, Label: "30天"},
	{Value: TtAwemeFanTime60Days, Label: "60天"},
}

// 媒体定向
const (
	TtSuperiorPopularityNone = "NONE"
	TtSuperiorPopularityGame = "GAME"
)

var TtSuperiorPopularityTypeOptions = []TtStringOption{
	{Value: TtSuperiorPopularityNone, Label: "不限"},
	{Value: TtSuperiorPopularityGame, Label: "游戏优质媒体"},
}

// 最低安卓版本
var TtAndroidOsvOptions = []TtStringOption{
	{Value: "NONE", Label: "不限"},
	{Value: "11.0", Label: "Android 11.0及以上"},
	{Value: "10.0", Label: "Android 10.0及以上"},
	{Value: "9.0", Label: "Android 9.0及以上"},
	{Value: "8.1", Label: "Android 8.1及以上"},
	{Value: "8.0", Label: "Android 8.0及以上"},
	{Value: "7.1", Label: "Android 7.1及以上"},
	{Value: "7.0", Label: "Android 7.0及以上"},
	{Value: "6.0", Label: "Android 6.0及以上"},
	{Value: "5.1", Label: "Android 5.1及以上"},
	{Value: "5.0", Label: "Android 5.0及以上"},
	{Value: "4.4", Label: "Android 4.4及以上"},
	{Value: "4.0", Label: "Android 4.0及以上"},
	{Value: "3.0", Label: "Android 3.0及以上"},
}

// 最低IOS版本
var TtIosOsvOptions = []TtStringOption{
	{Value: "NONE", Label: "不限"},
	{Value: "14.0", Label: "iOS 14.0及以上"},
	{Value: "13.0", Label: "iOS 13.0及以上"},
	{Value: "12.0", Label: "iOS 12.0及以上"},
	{Value: "11.0", Label: "iOS 11.0及以上"},
	{Value: "10.0", Label: "iOS 10.0及以上"},
	{Value: "9.0", Label: "iOS 9.0及以上"},
	{Value: "8.0", Label: "iOS 8.0及以上"},
	{Value: "7.0", Label: "iOS 7.0及以上"},
	{Value: "6.0", Label: "iOS 6.0及以上"},
	{Value: "5.0", Label: "iOS 5.0及以上"},
	{Value: "4.0", Label: "iOS 4.0及以上"},
}

// 鸿蒙版本
var TtHarmonyOsvOptions = []TtStringOption{
	{Value: "NONE", Label: "不限"},
	{Value: "5.0", Label: "5.0"},
}

// 设备类型
const (
	TtDeviceTypeMobile = "MOBILE"
	TtDeviceTypePad    = "PAD"
)

var TtDeviceTypeOptions = []TtStringOption{
	{Value: TtDeviceTypeMobile, Label: "手机"},
	{Value: TtDeviceTypePad, Label: "平板"},
}

// 网络类型
const (
	TtAcWifi = "WIFI"
	TtAc2G   = "2G"
	TtAc3G   = "3G"
	TtAc4G   = "4G"
	TtAc5G   = "5G"
)

var TtAcTypeOptions = []TtStringOption{
	{Value: TtAcWifi, Label: "WIFI"},
	{Value: TtAc2G, Label: "2G"},
	{Value: TtAc3G, Label: "3G"},
	{Value: TtAc4G, Label: "4G"},
	{Value: TtAc5G, Label: "5G"},
}

// 运营商
const (
	TtCarrierMobile = "MOBILE"
	TtCarrierUnicom = "UNICOM"
	TtCarrierTelcom = "TELCOM"
)

var TtCarrierOptions = []TtStringOption{
	{Value: TtCarrierMobile, Label: "移动"},
	{Value: TtCarrierUnicom, Label: "联通"},
	{Value: TtCarrierTelcom, Label: "电信"},
}

// 过滤是否安装
const (
	TtHideIfExistsUnlimited = "UNLIMITED"
	TtHideIfExistsFilter    = "FILTER"
	TtHideIfExistsTargeting = "TARGETING"
)

var TtHideIfExistsOptions = []TtStringOption{
	{Value: TtHideIfExistsUnlimited, Label: "不限"},
	{Value: TtHideIfExistsFilter, Label: "过滤"},
	{Value: TtHideIfExistsTargeting, Label: "定向"},
}

// 过滤已转化用户
const (
	TtHideIfConvertedNoExclude  = "NO_EXCLUDE"
	TtHideIfConvertedPromotion  = "PROMOTION"
	TtHideIfConvertedProject    = "PROJECT"
	TtHideIfConvertedAdvertiser = "ADVERTISER"
	TtHideIfConvertedApp        = "APP"
	TtHideIfConvertedCustomer   = "CUSTOMER"
	TtHideIfConvertedOrg        = "ORGANIZATION"
)

var TtHideIfConvertedOptions = []TtStringOption{
	{Value: TtHideIfConvertedNoExclude, Label: "不限"},
	{Value: TtHideIfConvertedPromotion, Label: "广告"},
	{Value: TtHideIfConvertedProject, Label: "推广项目"},
	{Value: TtHideIfConvertedAdvertiser, Label: "广告账户"},
	{Value: TtHideIfConvertedApp, Label: "APP"},
	{Value: TtHideIfConvertedCustomer, Label: "公司账户"},
	{Value: TtHideIfConvertedOrg, Label: "组织（仅加白广告主可用）"},
}

// 过滤时间范围
const (
	TtConvertTimeNone        = "NONE"
	TtConvertTimeToday       = "TODAY"
	TtConvertTimeSevenDay    = "SEVEN_DAY"
	TtConvertTimeOneMonth    = "ONE_MONTH"
	TtConvertTimeThreeMonth  = "THREE_MONTH"
	TtConvertTimeSixMonth    = "SIX_MONTH"
	TtConvertTimeTwelveMonth = "TWELVE_MONTH"
)

var TtConvertTimeDurationOptions = []TtStringOption{
	{Value: TtConvertTimeNone, Label: "不限"},
	{Value: TtConvertTimeToday, Label: "当天"},
	{Value: TtConvertTimeSevenDay, Label: "7天"},
	{Value: TtConvertTimeOneMonth, Label: "1个月"},
	{Value: TtConvertTimeThreeMonth, Label: "3个月"},
	{Value: TtConvertTimeSixMonth, Label: "6个月"},
	{Value: TtConvertTimeTwelveMonth, Label: "12个月"},
}

// 过滤高活跃用户 / 自己的粉丝 / 高关注数用户
const (
	TtFilterNone = "NONE"
	TtFilterOn   = "ON"
)

var TtFilterAwemeAbnormalActiveOptions = []TtStringOption{
	{Value: TtFilterNone, Label: "不过滤"},
	{Value: TtFilterOn, Label: "过滤"},
}

var TtFilterOwnAwemeFansOptions = []TtStringOption{
	{Value: TtFilterNone, Label: "不过滤"},
	{Value: TtFilterOn, Label: "过滤"},
}

var TtFilterAwemeFansCountOptions = []TtOption{
	{Value: 0, Label: "不过滤"},
	{Value: 200, Label: ">200"},
	{Value: 500, Label: ">500"},
	{Value: 1000, Label: ">1000"},
}

// 新用户使用头条时间范围
const (
	TtActivateUnlimited       = "UNLIMITED"
	TtActivateWithinAMonth    = "WITH_IN_A_MONTH"
	TtActivateOneToThreeMonth = "ONE_MONTH_2_THREE_MONTH"
	TtActivateThreeMonthEarly = "THREE_MONTH_EAILIER"
)

var TtActivateTypeOptions = []TtStringOption{
	{Value: TtActivateUnlimited, Label: "不限"},
	{Value: TtActivateWithinAMonth, Label: "一个月以内"},
	{Value: TtActivateOneToThreeMonth, Label: "一个月到三个月"},
	{Value: TtActivateThreeMonthEarly, Label: "三个月或更早"},
}

// 手机品牌
const (
	TtDeviceBrandHuawei  = "HUAWEI"
	TtDeviceBrandApple   = "APPLE"
	TtDeviceBrandHonor   = "HONOR"
	TtDeviceBrandXiaomi  = "XIAOMI"
	TtDeviceBrandSamsung = "SAMSUNG"
	TtDeviceBrandOppo    = "OPPO"
	TtDeviceBrandVivo    = "VIVO"
	TtDeviceBrandMeizu   = "MEIZU"
	TtDeviceBrandNokia   = "NOKIA"
	TtDeviceBrandOneplus = "ONEPLUS"
)

var TtDeviceBrandOptions = []TtStringOption{
	{Value: TtDeviceBrandHuawei, Label: "华为"},
	{Value: TtDeviceBrandApple, Label: "苹果"},
	{Value: TtDeviceBrandHonor, Label: "荣耀"},
	{Value: TtDeviceBrandXiaomi, Label: "小米"},
	{Value: TtDeviceBrandSamsung, Label: "三星"},
	{Value: TtDeviceBrandOppo, Label: "OPPO"},
	{Value: TtDeviceBrandVivo, Label: "VIVO"},
	{Value: TtDeviceBrandMeizu, Label: "魅族"},
	{Value: TtDeviceBrandNokia, Label: "诺基亚"},
	{Value: TtDeviceBrandOneplus, Label: "一加"},
}

// 智能放量
const (
	TtAutoExtendOn  = "ON"
	TtAutoExtendOff = "OFF"
)

var TtAutoExtendEnabledOptions = []TtStringOption{
	{Value: TtAutoExtendOn, Label: "开启"},
	{Value: TtAutoExtendOff, Label: "关闭"},
}

// 可放开定向
const (
	TtAutoExtendAge            = "AGE"
	TtAutoExtendRegion         = "REGION"
	TtAutoExtendGender         = "GENDER"
	TtAutoExtendCustomAudience = "CUSTOM_AUDIENCE"
)

var TtAutoExtendTargetOptions = []TtStringOption{
	{Value: TtAutoExtendAge, Label: "年龄"},
	{Value: TtAutoExtendRegion, Label: "地域"},
	{Value: TtAutoExtendGender, Label: "性别"},
	{Value: TtAutoExtendCustomAudience, Label: "自定人群-定向"},
}

// 投放平台（定向包）
const (
	TtPlatformIos     = "IOS"
	TtPlatformAndroid = "ANDROID"
	TtPlatformHarmony = "HARMONY"
)

var TtPlatformNameOptions = []TtStringOption{
	{Value: TtPlatformIos, Label: "IOS"},
	{Value: TtPlatformAndroid, Label: "安卓"},
	{Value: TtPlatformHarmony, Label: "鸿蒙"},
}

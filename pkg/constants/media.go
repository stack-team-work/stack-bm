package constants

// 渠道标识常量
const (
	ChannelTt    = "TT"       // 头条
	ChannelTc    = "TC"       // 腾讯
	ChannelBili  = "BILIBILI" // B站
	ChannelKs    = "KS"       // 快手
)

// MediaMarkChannel media.mark(小写唯一索引) -> 渠道标识
var MediaMarkChannel = map[string]string{
	"tt":   ChannelTt,
	"tc":   ChannelTc,
	"bili": ChannelBili,
	"ks":   ChannelKs,
}

// ChannelMark 渠道标识 -> media.mark，反向映射
var ChannelMark = map[string]string{
	ChannelTt:   "tt",
	ChannelTc:   "tc",
	ChannelBili: "bili",
	ChannelKs:   "ks",
}

// ChannelName 渠道标识 -> 中文名
var ChannelName = map[string]string{
	ChannelTt:   "头条",
	ChannelTc:   "腾讯",
	ChannelBili: "B站",
	ChannelKs:   "快手",
}

// 管家授权状态（对齐源 MktAccountManager::AUTH_STATUS）
const (
	ManagerAuthStatusWait        = 0 // 未授权
	ManagerAuthStatusComplete    = 1 // 已授权
	ManagerAuthStatusAuthFail    = 2 // 授权失败
	ManagerAuthStatusRefreshFail = 3 // 刷新授权失败
)

// 管家授权版本（兼容B站V1/V2）
const (
	ManagerVersionV1 = 1
	ManagerVersionV2 = 2
)

// OAuth 回调基础路径模板（%s 为渠道标识小写）
const OAuthRedirectPath = "%s"

// 广告数据层级（统一命名）
const (
	AdDataLevelAccount  = "account"  // 账户
	AdDataLevelCampaign = "campaign" // 第一层级：广告组
	AdDataLevelUnit     = "unit"     // 第二层级：广告
	AdDataLevelCreative = "creative" // 第三层级：创意
)

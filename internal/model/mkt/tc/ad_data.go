package tc

// 广告数据集合名（对齐源 mktAdData*StatementList.table）
// 腾讯只含 账户 / 第一层级(campaign, 源广告组) / 第二层级(unit, 源计划)
const (
	AdDataCollectionAccount  = "mktAdDataTcAccountStatementList.table"
	AdDataCollectionCampaign = "mktAdDataTcAdgroupStatementList.table"
	AdDataCollectionUnit     = "" // 腾讯暂无计划(unit)数据源集合，列表返回空
)

// AdDataLevelCollections 广告数据层级 -> 集合名
var AdDataLevelCollections = map[string]string{
	"account":  AdDataCollectionAccount,
	"campaign": AdDataCollectionCampaign,
	"unit":     AdDataCollectionUnit,
}
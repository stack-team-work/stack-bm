package tc

// 广告数据集合名（对齐源 mktAdData*StatementList.table）
const (
	AdDataCollectionAccount  = "mktAdDataTcAccountStatementList.table"
	AdDataCollectionCampaign = "mktAdDataTcAdgroupStatementList.table"
	AdDataCollectionCreative = "mktAdDataTcCreativeStatementList.table"
	AdDataCollectionUnit     = "" // 腾讯暂无广告(unit)数据源集合，列表返回空
)

// AdDataLevelCollections 广告数据层级 -> 集合名
var AdDataLevelCollections = map[string]string{
	"account":  AdDataCollectionAccount,
	"campaign": AdDataCollectionCampaign,
	"unit":     AdDataCollectionUnit,
	"creative": AdDataCollectionCreative,
}
package bili

// 广告数据集合名（对齐源 mktAdData*StatementList.table）
const (
	AdDataCollectionAccount  = "mktAdDataBilibiliAccountStatementList.table"
	AdDataCollectionCampaign = "mktAdDataBilibiliCampaignStatementList.table"
	AdDataCollectionUnit     = "mktAdDataBiliUnitStatementList.table"
	AdDataCollectionCreative = "mktAdDataBiliCreativeStatementList.table"
)

// AdDataLevelCollections 广告数据层级 -> 集合名
var AdDataLevelCollections = map[string]string{
	"account":  AdDataCollectionAccount,
	"campaign": AdDataCollectionCampaign,
	"unit":     AdDataCollectionUnit,
	"creative": AdDataCollectionCreative,
}
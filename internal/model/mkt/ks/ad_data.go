package ks

// 广告数据集合名（对齐源 mktAdData*StatementList.table）
const (
	AdDataCollectionAccount  = "mktAdDataKuaiShouAccountStatementList.table"
	AdDataCollectionCampaign = "mktAdDataKuaiShouCampaignStatementList.table"
	AdDataCollectionUnit     = "mktAdDataKuaiShouUnitStatementList.table"
	AdDataCollectionCreative = "mktAdDataKuaiShouCreativeStatementList.table"
)

// AdDataLevelCollections 广告数据层级 -> 集合名
var AdDataLevelCollections = map[string]string{
	"account":  AdDataCollectionAccount,
	"campaign": AdDataCollectionCampaign,
	"unit":     AdDataCollectionUnit,
	"creative": AdDataCollectionCreative,
}
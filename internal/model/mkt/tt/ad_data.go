package tt

// 广告数据集合名（对齐源 mktAdData*StatementList.table）
// 头条仅迁移 V3，只含 账户 / 第一层级(campaign, 源项目)
const (
	AdDataCollectionAccount  = "mktAdDataTtV3AccountStatementList.table"
	AdDataCollectionCampaign = "mktAdDataTtV3ProjectStatementList.table"
)

// AdDataLevelCollections 广告数据层级 -> 集合名
var AdDataLevelCollections = map[string]string{
	"account":  AdDataCollectionAccount,
	"campaign": AdDataCollectionCampaign,
}
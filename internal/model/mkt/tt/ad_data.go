package tt

// 广告数据集合名（对齐源 mktAdData*StatementList.table）
// 头条仅迁移 V3：账户/项目(映射campaign)/广告(映射unit)
const (
	AdDataCollectionAccount  = "mktAdDataTtV3AccountStatementList.table"
	AdDataCollectionCampaign = "mktAdDataTtV3ProjectStatementList.table"
	AdDataCollectionUnit     = "mktAdDataTtPromotionStatementList.table"
	AdDataCollectionCreative = "" // 头条V3暂无创意(creative)数据源集合，列表返回空
)

// AdDataLevelCollections 广告数据层级 -> 集合名
var AdDataLevelCollections = map[string]string{
	"account":  AdDataCollectionAccount,
	"campaign": AdDataCollectionCampaign,
	"unit":     AdDataCollectionUnit,
	"creative": AdDataCollectionCreative,
}
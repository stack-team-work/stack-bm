package ks

// AdDataRepository 广告数据仓储
// 当前表（MongoDB mktAdData*StatementList 集合）尚未创建，列表先返回空；
// 表建好后在此实现真实聚合查询（对齐源 MediaData 模块按 cpid 分组 + 报表指标聚合）。
type AdDataRepository struct{}

func NewAdDataRepository() *AdDataRepository { return &AdDataRepository{} }

// List 查询广告数据列表，当前返回空
func (r *AdDataRepository) List(collection string, page, size int, params map[string]interface{}) ([]map[string]interface{}, int64, error) {
	return []map[string]interface{}{}, 0, nil
}
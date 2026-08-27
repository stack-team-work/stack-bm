package ttad

import (
	ttModel "stack-bm/internal/model/mkt/tt"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// AccountRepository 头条账户数据仓储（集合尚未创建，List 先返回空）
type AccountRepository struct{}

func NewAccountRepository() *AccountRepository { return &AccountRepository{} }

func (r *AccountRepository) col() *mongo.Collection {
	return collection(ttModel.AdDataCollectionAccount)
}

// List 分页查询账户数据，columns 为字段投影，filters 为筛选条件
func (r *AccountRepository) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	return []map[string]interface{}{}, 0, nil
}

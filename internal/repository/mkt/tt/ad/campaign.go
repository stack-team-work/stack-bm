package ttad

import (
	"context"
	"errors"

	ttModel "stack-bm/internal/model/mkt/tt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CampaignRepository 头条V3项目（第一层级）数据仓储
type CampaignRepository struct{}

func NewCampaignRepository() *CampaignRepository { return &CampaignRepository{} }

const campaignIDField = "cpid"

func (r *CampaignRepository) col() *mongo.Collection {
	return collection(ttModel.AdDataCollectionCampaign)
}

// List 分页查询项目数据（集合建好后实现真实聚合查询）
func (r *CampaignRepository) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	return []map[string]interface{}{}, 0, nil
}

// Upsert 按 cpid upsert 同步数据
func (r *CampaignRepository) Upsert(doc bson.M) error {
	col := r.col()
	if col == nil {
		return errMongoUnavailable
	}
	_, err := col.UpdateOne(context.Background(),
		bson.M{campaignIDField: doc[campaignIDField]},
		bson.M{"$set": doc},
		options.UpdateOne().SetUpsert(true),
	)
	return err
}

// ResolveAccount 通过项目ID(cpid) 解析平台账户ID(account_id)
func (r *CampaignRepository) ResolveAccount(id int) (int, error) {
	col := r.col()
	if col == nil {
		return 0, errMongoUnavailable
	}
	var raw bson.M
	err := col.FindOne(context.Background(), bson.M{campaignIDField: id}).Decode(&raw)
	if err == mongo.ErrNoDocuments {
		return 0, errors.New("未找到该项目数据，无法解析渠道账户")
	}
	if err != nil {
		return 0, err
	}
	return accountIDFromDoc(raw)
}

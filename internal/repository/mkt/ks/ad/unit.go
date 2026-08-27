package ksad

import (
	"context"
	"errors"

	ksModel "stack-bm/internal/model/mkt/ks"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// UnitRepository 快手广告（第二层级）数据仓储
type UnitRepository struct{}

func NewUnitRepository() *UnitRepository { return &UnitRepository{} }

const unitIDField = "aid"

func (r *UnitRepository) col() *mongo.Collection {
	return collection(ksModel.AdDataCollectionUnit)
}

// List 分页查询广告数据
func (r *UnitRepository) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	return []map[string]interface{}{}, 0, nil
}

// Upsert 按 aid upsert 同步数据
func (r *UnitRepository) Upsert(doc bson.M) error {
	col := r.col()
	if col == nil {
		return errMongoUnavailable
	}
	_, err := col.UpdateOne(context.Background(),
		bson.M{unitIDField: doc[unitIDField]},
		bson.M{"$set": doc},
		options.UpdateOne().SetUpsert(true),
	)
	return err
}

// ResolveAccount 通过广告ID(aid) 解析平台账户ID(account_id)
func (r *UnitRepository) ResolveAccount(id int) (int, error) {
	col := r.col()
	if col == nil {
		return 0, errMongoUnavailable
	}
	var raw bson.M
	err := col.FindOne(context.Background(), bson.M{unitIDField: id}).Decode(&raw)
	if err == mongo.ErrNoDocuments {
		return 0, errors.New("未找到该广告数据，无法解析渠道账户")
	}
	if err != nil {
		return 0, err
	}
	return accountIDFromDoc(raw)
}

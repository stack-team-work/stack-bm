package ksad

import (
	"context"
	"errors"

	ksModel "stack-bm/internal/model/mkt/ks"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// CreativeRepository 快手创意（第三层级）数据仓储
type CreativeRepository struct{}

func NewCreativeRepository() *CreativeRepository { return &CreativeRepository{} }

const creativeIDField = "cid"

func (r *CreativeRepository) col() *mongo.Collection {
	return collection(ksModel.AdDataCollectionCreative)
}

// List 分页查询创意数据
func (r *CreativeRepository) List(page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	return []map[string]interface{}{}, 0, nil
}

// Upsert 按 cid upsert 同步数据
func (r *CreativeRepository) Upsert(doc bson.M) error {
	col := r.col()
	if col == nil {
		return errMongoUnavailable
	}
	_, err := col.UpdateOne(context.Background(),
		bson.M{creativeIDField: doc[creativeIDField]},
		bson.M{"$set": doc},
		options.UpdateOne().SetUpsert(true),
	)
	return err
}

// ResolveAccount 通过创意ID(cid) 解析平台账户ID(account_id)
func (r *CreativeRepository) ResolveAccount(id int) (int, error) {
	col := r.col()
	if col == nil {
		return 0, errMongoUnavailable
	}
	var raw bson.M
	err := col.FindOne(context.Background(), bson.M{creativeIDField: id}).Decode(&raw)
	if err == mongo.ErrNoDocuments {
		return 0, errors.New("未找到该创意数据，无法解析渠道账户")
	}
	if err != nil {
		return 0, err
	}
	return accountIDFromDoc(raw)
}

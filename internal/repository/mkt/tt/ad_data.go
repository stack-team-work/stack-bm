package tt

import (
	"context"
	"errors"

	"stack-bm/internal/database"
	ttModel "stack-bm/internal/model/mkt/tt"
	"stack-bm/pkg/constants"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// AdDataRepository 广告数据仓储
// 当前表（MongoDB mktAdData*StatementList 集合）尚未创建，列表先返回空；
// 表建好后在此实现真实聚合查询（对齐源 MediaData 模块按 cpid 分组 + 报表指标聚合）。
type AdDataRepository struct{}

func NewAdDataRepository() *AdDataRepository { return &AdDataRepository{} }

// List 查询广告数据列表，当前返回空；columns 为字段投影，filters 为筛选条件
func (r *AdDataRepository) List(collection string, page, size int, columns []string, filters map[string]interface{}) ([]map[string]interface{}, int64, error) {
	return []map[string]interface{}{}, 0, nil
}

// FindAccountByLevelID 通过操作ID(cpid/aid/cid) 查询渠道账户id(account_id)
func (r *AdDataRepository) FindAccountByLevelID(level string, id int) (int, error) {
	collectionName, ok := ttModel.AdDataLevelCollections[level]
	if !ok || collectionName == "" {
		return 0, errors.New("未知层级: " + level)
	}
	idField, ok := constants.AdDataLevelIDField[level]
	if !ok {
		return 0, errors.New("未知层级ID字段: " + level)
	}
	if database.ChannelDB == nil {
		return 0, errMongoUnavailable
	}
	var doc struct {
		AccountID int `bson:"account_id"`
	}
	err := database.ChannelDB.Collection(collectionName).
		FindOne(context.Background(), bson.M{idField: id}).
		Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return 0, errors.New("未找到该广告数据，无法解析渠道账户")
	}
	if err != nil {
		return 0, err
	}
	return doc.AccountID, nil
}

// Upsert 按 account_id + 层级ID字段 upsert 同步数据
func (r *AdDataRepository) Upsert(level string, idField string, doc bson.M) error {
	collectionName, ok := ttModel.AdDataLevelCollections[level]
	if !ok || collectionName == "" || database.ChannelDB == nil {
		return errMongoUnavailable
	}
	filter := bson.M{idField: doc[idField]}
	_, err := database.ChannelDB.Collection(collectionName).
		UpdateOne(context.Background(), filter, bson.M{"$set": doc}, options.UpdateOne().SetUpsert(true))
	return err
}
package media

import (
	"context"
	"errors"

	"stack-bm/internal/database"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const ManagerTokenCollection = "mkt_account_manager_token"

var errMongoUnavailable = errors.New("MongoDB 未连接")

type MediaManagerTokenRepository struct {
	col *mongo.Collection
}

func NewMediaManagerTokenRepository() *MediaManagerTokenRepository {
	var col *mongo.Collection
	if database.ChannelRawDB != nil {
		col = database.ChannelRawDB.Collection(ManagerTokenCollection)
	}
	return &MediaManagerTokenRepository{col: col}
}

// Upsert 写入/更新管家 token 文档
func (r *MediaManagerTokenRepository) Upsert(managerID int, doc bson.M) error {
	if r.col == nil {
		return errMongoUnavailable
	}
	_, err := r.col.UpdateOne(context.Background(),
		bson.M{"mkt_account_manager_id": managerID},
		bson.M{"$set": doc},
		options.UpdateOne().SetUpsert(true),
	)
	return err
}

// FindByManagerID 查询管家 token 文档
func (r *MediaManagerTokenRepository) FindByManagerID(managerID int) (bson.M, error) {
	if r.col == nil {
		return nil, errMongoUnavailable
	}
	var doc bson.M
	err := r.col.FindOne(context.Background(), bson.M{"mkt_account_manager_id": managerID}).Decode(&doc)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return doc, nil
}

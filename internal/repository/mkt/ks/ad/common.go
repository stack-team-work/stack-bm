package ksad

import (
	"errors"
	"strconv"

	"stack-bm/internal/database"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var errMongoUnavailable = errors.New("MongoDB 未连接")

var errAccountIDType = errors.New("account_id 字段类型异常")

// collection 获取渠道 Mongo 集合
func collection(name string) *mongo.Collection {
	if database.ChannelDB == nil {
		return nil
	}
	return database.ChannelDB.Collection(name)
}

// accountIDFromDoc 归一 doc.account_id 为 int：新同步统一写 int，
// 存量数据存在字符串/浮点形态
func accountIDFromDoc(doc bson.M) (int, error) {
	switch v := doc["account_id"].(type) {
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, errAccountIDType
	}
}

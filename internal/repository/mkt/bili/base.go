package bili

import (
	"errors"

	"stack-bm/internal/database"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var errMongoUnavailable = errors.New("MongoDB 未连接")

func collection(name string) *mongo.Collection {
	if database.ChannelDB == nil {
		return nil
	}
	return database.ChannelDB.Collection(name)
}

func buildPageFilter(keyword string) bson.M {
	filter := bson.M{"display": 1}
	if keyword != "" {
		filter["template_name"] = bson.M{"$regex": keyword, "$options": "i"}
	}
	return filter
}

func pageOptions(page, size int) *options.FindOptionsBuilder {
	return options.Find().SetSkip(int64((page - 1) * size)).SetLimit(int64(size)).
		SetSort(bson.D{{Key: "updated_at", Value: -1}})
}

func parseID(id string) (bson.ObjectID, error) {
	return bson.ObjectIDFromHex(id)
}

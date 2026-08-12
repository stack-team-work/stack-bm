package tt

import (
	"context"

	"stack-bm/internal/database"
	ttModel "stack-bm/internal/model/mkt/tt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type WordListRepository struct{}

func NewWordListRepository() *WordListRepository { return &WordListRepository{} }

func (r *WordListRepository) ListAll() ([]ttModel.WordList, error) {
	if database.ChannelRawDB == nil {
		return nil, errMongoUnavailable
	}
	ctx := context.Background()
	cursor, err := rawCollection(ttModel.WordList{}.CollectionName()).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var list []ttModel.WordList
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (r *WordListRepository) ListByNames(names []string) ([]ttModel.WordList, error) {
	if database.ChannelRawDB == nil {
		return nil, errMongoUnavailable
	}
	if len(names) == 0 {
		return nil, nil
	}
	ctx := context.Background()
	cursor, err := rawCollection(ttModel.WordList{}.CollectionName()).Find(ctx, bson.M{"name": bson.M{"$in": names}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var list []ttModel.WordList
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}

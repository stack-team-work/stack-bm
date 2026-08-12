package bili

import (
	"context"

	"stack-bm/internal/database"
	biliModel "stack-bm/internal/model/mkt/bili"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AudienceTemplateRepository struct{}

func NewAudienceTemplateRepository() *AudienceTemplateRepository { return &AudienceTemplateRepository{} }

func (r *AudienceTemplateRepository) col() *mongo.Collection { return collection("bili_audience_template") }

func (r *AudienceTemplateRepository) Create(doc *biliModel.AudienceTemplate) error {
	if database.ChannelDB == nil {
		return errMongoUnavailable
	}
	res, err := r.col().InsertOne(context.Background(), doc)
	if err != nil {
		return err
	}
	doc.ID = res.InsertedID.(bson.ObjectID)
	return nil
}

func (r *AudienceTemplateRepository) FindByID(id string) (*biliModel.AudienceTemplate, error) {
	if database.ChannelDB == nil {
		return nil, errMongoUnavailable
	}
	oid, err := parseID(id)
	if err != nil {
		return nil, err
	}
	var doc biliModel.AudienceTemplate
	if err := r.col().FindOne(context.Background(), bson.M{"_id": oid}).Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

func (r *AudienceTemplateRepository) FindPage(page, size int, keyword string) ([]biliModel.AudienceTemplate, int64, error) {
	if database.ChannelDB == nil {
		return nil, 0, errMongoUnavailable
	}
	ctx := context.Background()
	filter := buildPageFilter(keyword)
	total, err := r.col().CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	cursor, err := r.col().Find(ctx, filter, pageOptions(page, size))
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)
	var list []biliModel.AudienceTemplate
	if err := cursor.All(ctx, &list); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *AudienceTemplateRepository) Update(doc *biliModel.AudienceTemplate) error {
	if database.ChannelDB == nil {
		return errMongoUnavailable
	}
	_, err := r.col().ReplaceOne(context.Background(), bson.M{"_id": doc.ID}, doc)
	return err
}

func (r *AudienceTemplateRepository) SoftDelete(id string) error {
	if database.ChannelDB == nil {
		return errMongoUnavailable
	}
	oid, err := parseID(id)
	if err != nil {
		return err
	}
	_, err = r.col().UpdateOne(context.Background(), bson.M{"_id": oid},
		bson.M{"$set": bson.M{"display": biliModel.DisplayHide}})
	return err
}

func (r *AudienceTemplateRepository) ListAll() ([]biliModel.AudienceTemplate, error) {
	if database.ChannelDB == nil {
		return nil, errMongoUnavailable
	}
	ctx := context.Background()
	cursor, err := r.col().Find(ctx, bson.M{"display": 1},
		options.Find().SetSort(bson.M{"updated_at": -1}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var list []biliModel.AudienceTemplate
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}

package events

import (
	"context"
	"log"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
	"github.com/ibra-bybuy/go-wsports-events/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoRepository) GetByQuery(ctx context.Context, query string, limit, page int) *model.PaginationResponse {
	response := model.PaginationResponse{
		Items:      []model.Event{},
		Pagination: model.Pagination{},
	}

	collection, ctx, cancel := m.getCollection()
	defer cancel()

	model := mongo.IndexModel{Keys: bson.D{{Key: "name", Value: "text"}}}
	_, err := collection.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		log.Println(err)
		return &response
	}

	findFilter := bson.D{{Key: "$text", Value: bson.D{{Key: "$search", Value: "ufc"}}}, {Key: "startAt", Value: bson.D{{Key: "$gt", Value: utils.StartDate("")}}}}
	// COUNT
	totalItems, err := collection.CountDocuments(context.TODO(), findFilter)
	if err != nil {
		return &response
	}
	response.Pagination = *utils.BuildPagination(totalItems, page, limit)

	// FIND
	opts := options.Find().SetSort(bson.D{{Key: "startAt", Value: 1}})
	l := int64(limit)
	skip := int64(page*limit - limit)
	opts.Limit = &l
	opts.Skip = &skip
	cursor, err := collection.Find(ctx, findFilter, opts)

	if err != nil {
		log.Println(err)
		return &response
	}

	err = cursor.All(ctx, &response.Items)

	return &response
}

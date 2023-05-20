package events

import (
	"context"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
	"github.com/ibra-bybuy/go-wsports-events/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoRepository) GetByTournament(ctx context.Context, tournament string, limit, page int) *model.PaginationResponse {
	response := model.PaginationResponse{
		Items:      []model.Event{},
		Pagination: model.Pagination{},
	}

	collection, ctx, cancel := m.getCollection()
	defer cancel()

	findFilter := bson.D{{Key: "name", Value: tournament}}

	// COUNT
	countOpts := options.Count().SetHint("_id_")
	totalItems, err := collection.CountDocuments(context.TODO(), findFilter, countOpts)
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
		return &response
	}

	err = cursor.All(ctx, &response.Items)

	return &response
}

package events

import (
	"context"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) GetByID(ctx context.Context, ID string) (*model.Event, error) {
	response := model.Event{}
	collection, ctx, cancel := m.getCollection()
	defer cancel()

	findFilter := bson.D{{Key: "id", Value: ID}}

	result := collection.FindOne(ctx, findFilter)

	if result.Err() != nil {
		return nil, errNotFound
	}

	result.Decode(&response)

	return &response, nil
}

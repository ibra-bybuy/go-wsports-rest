package events

import (
	"context"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *MongoRepository) GetByUuid(ctx context.Context, uuid string) (*model.Event, error) {
	response := model.Event{}
	collection, ctx, cancel := m.getCollection()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return nil, errNotFound
	}

	findFilter := bson.D{{Key: "_id", Value: objectId}}

	result := collection.FindOne(ctx, findFilter)

	if result.Err() != nil {
		return nil, errNotFound
	}

	result.Decode(&response)

	return &response, nil
}

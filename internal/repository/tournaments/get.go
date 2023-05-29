package events

import (
	"context"
	"fmt"

	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
	"github.com/ibra-bybuy/go-wsports-events/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoRepository) Get(ctx context.Context) *[]model.Tournament {
	response := []model.Tournament{}

	collection, ctx, cancel := m.getCollection()
	defer cancel()

	endDateStr := utils.EndDate()

	// FIND
	findFilter := bson.D{{Key: "startAt", Value: bson.D{{Key: "$gte", Value: endDateStr}}}}
	values, err := collection.Distinct(ctx, "name", findFilter)

	if err != nil {
		return &response
	}
	for _, event := range values {
		response = append(response, model.Tournament{
			Code: fmt.Sprintf("%s", event),
			Name: fmt.Sprintf("%s", event),
		})
	}

	return &response
}

package events

import (
	"context"
	"time"

	"github.com/ibra-bybuy/go-wsports-events/internal/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	client *mongodb.Client
}

func New(client *mongodb.Client) *MongoRepository {
	return &MongoRepository{client}
}

func (m *MongoRepository) getCollection() (*mongo.Collection, context.Context, context.CancelFunc) {
	collection := m.client.Database("mongo").Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return collection, ctx, cancel
}

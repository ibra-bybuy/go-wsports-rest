package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Event struct {
	Uuid         primitive.ObjectID `json:"uuid" bson:"_id,omitempty"`
	ID           string             `json:"id" bson:"id"`
	Name         string             `json:"name" bson:"name"`
	Teams        []Team             `json:"teams" bson:"teams"`
	StartAt      string             `json:"startAt" bson:"startAt"`
	EndAt        string             `json:"endAt" bson:"endAt"`
	AvatarURL    string             `json:"avatarUrl" bson:"avatarUrl"`
	Address      string             `json:"address" bson:"address"`
	Lang         Lang               `json:"lang" bson:"lang"`
	Streams      []Stream           `json:"streams" bson:"streams"`
	HideElements string             `json:"hideElements" bson:"hideElements"`
	Sport        string             `json:"sport" bson:"sport"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Config struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Key   string             `bson:"key" json:"key"`
	Value float64            `bson:"value" json:"value"`
}

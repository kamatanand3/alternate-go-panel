package repositories

import (
	"context"
	"time"

	"altpanel/config"
	"altpanel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var configCollection *mongo.Collection

// Init function to set collection after DB connection
func InitConfigRepository() {
	configCollection = config.GetCollection("config")
}



func GetAllConfig() ([]models.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := configCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var configs []models.Config
	for cursor.Next(ctx) {
		var user models.Config
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		configs = append(configs, user)
	}
	return configs, nil
}


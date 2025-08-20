package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found, using system environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB")

	if mongoURI == "" || dbName == "" {
		log.Fatal("❌ MONGO_URI or MONGO_DB is not set in .env")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database(dbName)
	fmt.Println("✅ MongoDB connected to", dbName)
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}

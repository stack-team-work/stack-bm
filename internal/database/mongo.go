package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"stack-bm/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Mongo     *mongo.Client
	ChannelDB *mongo.Database
)

func InitMongo() {
	if config.AppConfig.Mongo.URI == "" {
		log.Println("MongoDB: MONGO_URI not configured, skip")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(config.AppConfig.Mongo.URI))
	if err != nil {
		log.Printf("MongoDB connect failed: %v", err)
		return
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Printf("MongoDB ping failed: %v", err)
		return
	}
	Mongo = client
	dbName := config.AppConfig.Mongo.ChannelDB
	if dbName == "" {
		dbName = "channel_template"
	}
	ChannelDB = client.Database(dbName)
	fmt.Printf("Connected to MongoDB: %s (%s)\n", config.AppConfig.Mongo.URI, dbName)
}

package utils

import (
	"context"
	"fmt"
	"log"
	"sgs_fullstack_challenge/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(config configs.Config) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.MongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to DB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error pinging the DB: %v", err)
	}

	return client.Database(config.MongoDatabaseName), nil
}

func IsCollectionEmpty(collection *mongo.Collection) bool {
	count, err := collection.CountDocuments(context.Background(), map[string]interface{}{})

	if err != nil {
		log.Fatalf("Failed to get documents count: %v", err)
		return false
	}

	return count == 0
}

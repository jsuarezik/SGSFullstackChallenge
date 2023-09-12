package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoUri            string
	MongoDatabaseName   string
	MongoCollectionName string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	return Config{
		MongoUri:            os.Getenv("MONGO_URI"),
		MongoDatabaseName:   os.Getenv("MONGO_DATABASE_NAME"),
		MongoCollectionName: os.Getenv("MONGO_COLLECTION_NAME"),
	}
}

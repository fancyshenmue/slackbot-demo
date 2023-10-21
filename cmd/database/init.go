package database

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	customDatabase "slackbot-demo/pkg/database"
)

var (
	MongoClient *mongo.Client
)

func init() {
	// Connect MongoDB
	conn := customDatabase.ConnectMongoDB(
		os.Getenv("_MONGO_USERNAME"),
		os.Getenv("_MONGO_PASSWORD"),
		os.Getenv("_MONGO_ENDPOINT"),
		os.Getenv("_MONGO_PORT"),
	)
	MongoClient = conn.ConnectDatabase()
}

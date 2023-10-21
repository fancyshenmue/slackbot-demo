package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDBDatabase              = "slackBot"
	MongoDBCollectionBotOperator = "botOperator"
)

type MongoOperatorStruct struct {
	Client     *mongo.Client
	Collection string
}

type DatabaseOperatorConnectStruct struct {
	UserName string
	Password string
	Endpoint string
	Port     string
}

func (d *DatabaseOperatorConnectStruct) ConnectDatabase() *mongo.Client {
	var uri = "mongodb://" + d.UserName + ":" + d.Password + "@" + d.Endpoint + ":" + d.Port
	opts := options.Client().ApplyURI(uri)

	// create client
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	return client
}

func ConnectMongoDB(username, password, endpoint, port string) *DatabaseOperatorConnectStruct {
	client := DatabaseOperatorConnectStruct{
		UserName: username,
		Password: password,
		Endpoint: endpoint,
		Port:     port,
	}
	return &client
}

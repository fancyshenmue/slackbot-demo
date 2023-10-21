package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataInsertStruct struct {
	Client     *mongo.Client
	Database   string
	Collection string
	Field      string
	Data       string
	Filter     bson.M
}

func (t *DataInsertStruct) DataInsert() {
	// debug
	fmt.Println("--------------------------------------------------")
	fmt.Printf("\tUpdate Field: %s\n", t.Field)
	fmt.Println("--------------------------------------------------")
	// insert documents
	coll := t.Client.Database(t.Database).Collection(t.Collection)
	opts := options.Update().SetUpsert(true)
	updateField := make(bson.M)
	updateField[t.Field] = t.Data
	documents := bson.M{
		"$set": updateField,
	}
	_, err := coll.UpdateOne(context.TODO(), t.Filter, documents, opts)
	if err != nil {
		log.Fatalf("Insert Documents Failed: %v", err)
	}
}

package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID string `json:"id"`
	Title string `json:"title"`
}

type Filter struct {
	ID string `bson:"id,omitempty"`
}

func getRepository() mongo.Collection {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(context, options.Client().ApplyURI("mongodb://localhost:27017"))
	if (err != nil) {
		log.Fatalln(err)
	}
	collection := client.Database("testapi").Collection("products")
	return *collection
}

func Find(filter Filter) []bson.M {
	collection := getRepository()
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	doc, _ := collection.Find(context, filter)
	var result []bson.M
	doc.All(context, &result)
	return result
}

func Insert(document Product) *mongo.InsertOneResult {
	collection := getRepository()
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := collection.InsertOne(context, document)
	if (err != nil) {
		log.Fatalln(err)
	}
	return response
}

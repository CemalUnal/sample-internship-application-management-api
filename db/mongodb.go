package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func Connect(address string) *mongo.Client {
	log.Println("Trying to connect to MongoDB at", address)
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("Couldn't connect to the MongoDB at adress %s. Error is: %s", address, err)
	} else {
		log.Println("Connected to MongoDB at", address)
	}

	return client
}

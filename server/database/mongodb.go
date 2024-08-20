package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func InitMongoDB(uri string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	db := client.Database("my_database")

	// Create a unique index on the email field in the users collection
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = db.Collection("users").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create unique index on email: %v", err)
	}

	return db
}

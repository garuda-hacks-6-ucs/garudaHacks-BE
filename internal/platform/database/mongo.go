package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoConnection creates a new connection to MongoDB and returns the database client.
func NewMongoConnection(ctx context.Context, uri, dbName string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// Set a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Ping the primary
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB!")
	return client.Database(dbName), nil
}

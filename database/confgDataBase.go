package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Context struct {
	CTX context.Context
}

func (context Context) CreateConnection() *mongo.Client {
	DATABASE_URL := os.Getenv("DATABASE_URL")

	clientOptions := options.Client().ApplyURI(DATABASE_URL)
	client, err := mongo.Connect(context.CTX, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

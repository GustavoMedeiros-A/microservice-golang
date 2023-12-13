package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001" // Not used yet
	mongoURL = "mongodb://mongo:27017"
	gRocPort = "50001" // Not used yet
)

var client *mongo.Client

type Config struct {
}

func main() {
	// connect to MongoDB
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	// create a context in order to disconnect
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)

	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	connection, err := mongo.Connect(context.TODO())
	if err != nil {
		log.Println("Error connecting to MongoDB: ", err)
		return nil, err
	}

	return connection, nil
}

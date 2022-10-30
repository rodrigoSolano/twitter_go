package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection is a variable to connect to the database
var MongoConnection = ConnectToDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://rodrigo:solano@cluster0.xevwb62.mongodb.net/?retryWrites=true&w=majority")

// ConnectToDB is a function to connect to the database
func ConnectToDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connected to MongoDB!")

	return client
}

// CheckConnection is a function to verify the connection to the database
func CheckConnection() bool {
	err := MongoConnection.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}

	return true
}

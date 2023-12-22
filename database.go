package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Your MongoDB Atlas Connection String
const uri = "mongodb+srv://masha:SFFLG4dKrInyBB2u@cocacolafria.ksukw21.mongodb.net/?retryWrites=true&w=majority"

// A global variable that will hold a reference to the MongoDB client
var mongoClient *mongo.Client

// The init function will run before our main function to establish a connection to MongoDB. If it cannot connect it will fail and the program will exit.
func init() {
	if err := connect_to_mongodb(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

func connect_to_mongodb() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}

// GET /movies - Get all movies
func getAnuncios(w http.ResponseWriter, r *http.Request) {
	// Find movies
	collection := mongoClient.Database("cocacolafria").Collection("anuncio")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Map results
	var anuncios []bson.M
	if err = cursor.All(context.Background(), &anuncios); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return movies
	json.NewEncoder(w).Encode(anuncios)
}

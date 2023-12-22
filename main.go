package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
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

func main() {
	r := chi.NewRouter()

	r.Get("/view/{title}", makeHandler(viewHandler))
	r.Get("/edit/{title}", makeHandler(editHandler))
	r.Post("/save/{title}", makeHandler(saveHandler))
	fileServer(r, "/static", http.Dir("static"))
	r.Get("/", landingPageHandler)
	r.Get("/anuncios", anunciosHandler)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	log.Printf("Listening on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

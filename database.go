package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

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

func fetchAnuncios() ([]bson.M, error) {
	// Find movies
	collection := mongoClient.Database("cocacolafria").Collection("anuncio")

	// Create a descending sort by _id
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}}).SetLimit(12)

	cursor, err := collection.Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	// Map results
	var anuncios []bson.M
	if err = cursor.All(context.Background(), &anuncios); err != nil {
		return nil, err
	}

	return anuncios, nil
}

func getAnuncios(w http.ResponseWriter, r *http.Request) {
	anuncios, err := fetchAnuncios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return movies
	json.NewEncoder(w).Encode(anuncios)
}

func getAnunciosHtml(w http.ResponseWriter, r *http.Request) {
	anuncios, err := fetchAnuncios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Format each anuncio into HTML

	for _, anuncio := range anuncios {
		title, ok := anuncio["title"].(string)
		if !ok {
			http.Error(w, "Error: anuncio title is not a string", http.StatusInternalServerError)
			return
		}

		description, ok := anuncio["description"].(string)
		if !ok {
			http.Error(w, "Error: anuncio description is not a string", http.StatusInternalServerError)
			return
		}

		// Start the response for each anuncio
		fmt.Fprint(w, `<div class="w-full sm:w-1/2 md:w-1/3 p-3">`)
		fmt.Fprint(w, `<div class="component rounded shadow p-5" _="on click add .clicked to me">`)
		fmt.Fprintf(w, `<div class="header text-xl">%s</div><div class="pt-5"><div class="messageContent">%s</div></div>`, title, description)
		// End the response for each anuncio
		fmt.Fprint(w, `</div></div>`)
	}
}

func createAnuncio(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		//Print the error message in the server log
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the values from the form data
	title := r.FormValue("title")
	description := r.FormValue("description")

	// Check that the fields are not empty
	if title == "" || description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<script>
				if (!document.querySelector('#errorMessages')) {
					var form = document.querySelector('#anuncioForm');
					var errorDiv = document.createElement('div');
					errorDiv.id = 'errorMessages';
					errorDiv.innerHTML = '<p style="color:red;">Debes ingresar título y descripción o Masha te enfría</p>';
					form.insertAfter(errorDiv, form.lastChild);
				}
			</script>
		`)
		return
	}

	if len(title) > 30 || len(description) > 250 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return
	}

	// Create a new anuncio
	anuncio := bson.M{
		"title":       title,
		"description": description,
		"date":        time.Now(),
		"game":        "cocacolafria",
		"location":    "Casa de Masha",
		"contact":     "gmail@mashamail.com",
	}

	// Insert the anuncio into the database
	collection := mongoClient.Database("cocacolafria").Collection("anuncio")
	_, err = collection.InsertOne(context.Background(), anuncio)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<script>
				if (!document.querySelector('#errorMessages')) {
					var form = document.querySelector('#anuncioForm');
					var errorDiv = document.createElement('div');
					errorDiv.id = 'errorMessages';
					errorDiv.innerHTML = '<p style="color:red;">Internal Server Error</p>';
					form.insertBefore(errorDiv, form.firstChild);
				}
			</script>
		`)
		return
	}

	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(`<script>clearAnuncioForm();window.location.reload();</script>`))
}

// ---- Options ----
func fetchOpciones() ([]bson.M, error) {
	// Find movies
	collection := mongoClient.Database("cocacolafria").Collection("opcion")

	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	// Map results
	var opciones []bson.M
	if err = cursor.All(context.Background(), &opciones); err != nil {
		return nil, err
	}

	return opciones, nil
}

func getOpcionesHtml(w http.ResponseWriter, r *http.Request) {
	opciones, err := fetchOpciones()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Format each opcion into HTML

	for _, opcion := range opciones {
		title, ok := opcion["title"].(string)
		if !ok {
			http.Error(w, "Error: opcion title is not a string", http.StatusInternalServerError)
			return
		}

		// count, ok := opcion["count"].(int32)
		// if !ok {opcion
		// 	http.Error(w, "Error: opcion count is not an int32", http.StatusInternalServerError)
		// 	return
		// }

		// Start the response for each opcion
		// fmt.Fprint(w, `<div class="w-full sm:w-1/2 md:w-1/3 p-3">`)
		fmt.Fprintf(w, `<button class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded">%s</button>`, title)
		// End the response for each opcion
		// fmt.Fprint(w, `</div>`)
	}
}

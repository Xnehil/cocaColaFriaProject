package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/view/{title}", makeHandler(viewHandler))
	r.Get("/edit/{title}", makeHandler(editHandler))
	r.Post("/save/{title}", makeHandler(saveHandler))
	fileServer(r, "/static", http.Dir("static"))
	r.Get("/", landingPageHandler)
	r.Get("/anuncios", anunciosHandler)
	r.Get("/getAnuncios", getAnuncios)

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

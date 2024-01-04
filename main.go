package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the Authorization header
		token := r.Header.Get("Authorization")

		// Validate the token (this is just a simple example, you should replace this with your actual validation logic)
		if token != "fantadepina" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the token is valid, call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()

	r.Get("/view/{title}", makeHandler(viewHandler))
	r.Get("/edit/{title}", makeHandler(editHandler))
	r.Post("/save/{title}", makeHandler(saveHandler))
	fileServer(r, "/static", http.Dir("static"))
	fileServer(r, "/templates", http.Dir("templates"))
	fileServer(r, "/scripts", http.Dir("scripts"))
	r.Get("/", landingPageHandler)
	r.Get("/anuncios", anunciosHandler)
	r.Get("/senado", senadoHandler)

	// Protected routes
	r.Route("/api", func(r chi.Router) {
		r.Use(AuthMiddleware) // Apply the AuthMiddleware to all routes in this group

		r.Get("/getAnuncios", getAnuncios)
		r.Get("/getAnunciosHtml", getAnunciosHtml)
		r.Get("/getVotacionesHtml", getVotacionesHtml)
		r.Post("/createAnuncio", createAnuncio)
		// Other protected routes...
	})

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

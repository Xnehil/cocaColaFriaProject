package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/go-chi/chi"
)

type Page struct {
	Title string

	Body []byte
}

func (p *Page) save() error {

	filename := p.Title + ".txt"

	return os.WriteFile(filename, p.Body, 0600)

}

func loadPage(title string) (*Page, error) {

	filename := title + ".txt"

	body, err := os.ReadFile(filename)

	if err != nil {

		return nil, err

	}

	return &Page{Title: title, Body: body}, nil

}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		http.Redirect(w, r, "/edit/"+title, http.StatusFound)

		return

	}

	renderTemplate(w, "view", p)

}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)

	if err != nil {

		p = &Page{Title: title}

	}

	renderTemplate(w, "edit", p)

}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")

	p := &Page{Title: title, Body: []byte(body)}

	err := p.save()

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return

	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)

}

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html", "templates/landing.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	err := templates.ExecuteTemplate(w, tmpl+".html", p)

	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {

			http.NotFound(w, r)

			return

		}

		fn(w, r, m[2])

	}

}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "Welcome"} // You can customize this Page struct as needed
	renderTemplate(w, "landing", p)
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func main() {
	r := chi.NewRouter()

	r.Get("/view/{title}", makeHandler(viewHandler))
	r.Get("/edit/{title}", makeHandler(editHandler))
	r.Post("/save/{title}", makeHandler(saveHandler))
	fileServer(r, "/static", http.Dir("static"))
	r.Get("/", landingPageHandler)

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

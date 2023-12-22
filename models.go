package main

type Page struct {
	Title string

	Body []byte
}

type Anuncio struct {
	Id          string
	Titulo      string
	Descripcion string
	Juego       string
	Fecha       string
	Lugar       string
	Contacto    string
}

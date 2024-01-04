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

type Opcion struct {
	Titulo string `bson:"titulo"`
	Cuenta int    `bson:"cuenta"`
}

type Votacion struct {
	Id          string   `bson:"_id"`
	Nombre      string   `bson:"nombre"`
	Descripcion string   `bson:"descripcion"`
	Activa      bool     `bson:"activa"`
	Opciones    []Opcion `bson:"opciones"`
}

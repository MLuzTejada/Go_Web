package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo")
}

func main() {

	http.HandleFunc("/", Hola)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

//serverMux: listado de rutas con acciones asociadas
//handler: acciones asociadas a las rutas, como respondo al cliente

package main

import (
	"fmt"
	"log"
	"net/http"
)

func hola(w http.ResponseWriter, r *http.Request) { //el identificador de la función se pone en mayuscula para usarlo desde otros paquetes
	fmt.Fprintf(w, "Hola Mundo")
}

func holaDos(w http.ResponseWriter, r *http.Request) { //el identificador de la función se pone en mayuscula para usarlo desde otros paquetes
	fmt.Fprintf(w, "Hola Mundo dos")
}

func main() {

	redirect := http.RedirectHandler("www.google.com.ar", http.StatusMovedPermanently)
	notfound := http.NotFoundHandler()

	mux := http.NewServeMux()
	mux.HandleFunc("/", hola)
	mux.HandleFunc("/dos", holaDos)
	mux.Handle("/redirect", redirect)
	mux.Handle("/not", notfound)

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}

//serverMux: listado de rutas con acciones asociadas
//handler: acciones asociadas a las rutas, como respondo al cliente

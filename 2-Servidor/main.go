package main

import (
	"fmt"
	"log"
	"net/http"
)

func hola(w http.ResponseWriter, r *http.Request) { //el identificador de la función se pone en mayuscula para usarlo desde otros paquetes
	fmt.Fprintf(w, "Hola Mundo")
}

func main() {

	http.HandleFunc("/", hola)
	log.Fatal(http.ListenAndServe("localhost:3000", nil)) //como esta nil, se utiliza DefaultServerMux
}

//serverMux: listado de rutas con acciones asociadas
//handler: acciones asociadas a las rutas, como respondo al cliente

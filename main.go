package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor header")
		fmt.Fprintf(w, "Hola Mundo")
	})

	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo, dos")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

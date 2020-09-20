package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Nombre", "Valor header")
		http.Redirect(w, r, "https://www.facebook.com/", http.StatusMovedPermanently)
	})

	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

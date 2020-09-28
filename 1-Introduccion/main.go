package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.URL.RawQuery)
		//fmt.Println(r.URL.Query())

		name := r.URL.Query().Get("name")
		if len(name) != 0 {
			fmt.Println("El nombre es:", name)
		}

		age := r.URL.Query().Get("age")
		if len(age) != 0 {
			fmt.Println("La edad es:", age)
		}

		fmt.Fprintf(w, "Hola Mundo")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

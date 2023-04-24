package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/name", nameHandler)
	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 page not found")
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Please provide a name in the query string, e.g. /name?name=Poe")
	} else {
		fmt.Fprintf(w, "Slt, %s!", name)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

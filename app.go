package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func BaseEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API : /hello")
}
func HelloWorldEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World - Go rest api")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", BaseEndPoint).Methods("GET")
	r.HandleFunc("/hello", HelloWorldEndPoint).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

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
	fmt.Fprintln(w, "pong")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", BaseEndPoint).Methods("GET")
	router.HandleFunc("/ping", HelloWorldEndPoint).Methods("GET")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

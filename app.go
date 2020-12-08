package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kushalshit27/go-rest-api/routers"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API : /hello")
}
func pingCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func main() {
	router := mux.NewRouter()
	const port string = ":8080"

	router.HandleFunc("/", hello).Methods("GET")
	router.HandleFunc("/ping", pingCheck).Methods("GET")

	router.HandleFunc("/posts", routers.GetPosts).Methods("GET")
	router.HandleFunc("/posts", routers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", routers.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", routers.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", routers.DeletePost).Methods("DELETE")

	log.Println("Server listening on port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}

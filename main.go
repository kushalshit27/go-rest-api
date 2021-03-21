package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kushalshit27/go-rest-api/internal/config"
	"github.com/kushalshit27/go-rest-api/internal/database"
	"github.com/kushalshit27/go-rest-api/internal/middleware"
	"github.com/kushalshit27/go-rest-api/internal/utils"
	"github.com/kushalshit27/go-rest-api/post"

	"github.com/gorilla/mux"
)

func routes(db *database.DB) *mux.Router {
	router := mux.NewRouter()

	postService := post.PostAPI(db)

	postService.Register(
		router.PathPrefix("/api").Subrouter(),
	)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(utils.ResponseError("error 404"))
	})

	return router
}

func main() {
	// Load config from the path provided by the user
	envPath := flag.String("env", ".env", "Path to environment file")
	flag.Parse()
	config := config.Load(*envPath)

	// Init database connection
	database := database.NewDB(config)
	defer database.Close()

	// Init server
	addr := fmt.Sprintf(":%s", "8080")
	log.Printf("[API] API running on: http://127.0.0.1:%s", "8080")

	server := http.Server{
		Addr:         addr,
		Handler:      middleware.CORS(routes(database)),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatalln(server.ListenAndServe())
}

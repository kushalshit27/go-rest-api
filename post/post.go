package post

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/kushalshit27/go-rest-api/internal/database"
	"github.com/kushalshit27/go-rest-api/internal/models"
	"github.com/kushalshit27/go-rest-api/internal/utils"
)

type postAPI struct {
	db *database.DB
}

func newPostAPI(db *database.DB) *postAPI {
	var postapi *postAPI
	postapi = new(postAPI)
	postapi.db = db
	return postapi
}

// All all
func (h *postAPI) All(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	rows, err := h.db.Query(ctx, "SELECT * FROM posts ORDER BY id ASC")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()

	var results []models.Post
	for rows.Next() {
		var r models.Post
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to scan %v\n", err)
			os.Exit(1)
		}
		results = append(results, r)
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess("", results))
}

// Get get
func (h *postAPI) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", "TODO"))
}

// Store store
func (h *postAPI) Store(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", "TODO"))
}

// Update update
func (h *postAPI) Update(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", "TODO"))
}

// Remove remove
func (h *postAPI) Remove(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", "TODO"))
}

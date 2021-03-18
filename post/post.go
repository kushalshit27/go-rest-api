package post

import (
	"net/http"
	"encoding/json"
	"github.com/kushalshit27/go-rest-api/internal/utils"
)

type postAPI struct{}

func newPostAPI() *postAPI {
	return new(postAPI)
}

// All all
func (h *postAPI) All(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", "TODO"))
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
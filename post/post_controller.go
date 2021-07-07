package post

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kushalshit27/go-rest-api/internal/models"
	"github.com/kushalshit27/go-rest-api/internal/utils"
)

type PostController interface {
	All(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	Store(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Remove(response http.ResponseWriter, request *http.Request)
}

type controller struct{}

var (
	postService PostService
)

func NewPostController(service PostService) PostController {
	postService = service
	return &controller{}

}

// All all
func (h *controller) All(w http.ResponseWriter, r *http.Request) {
	posts, err := postService.FindAll()
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError("Error getting the posts"))
		return
	}
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", posts))
}

// Get get
func (h *controller) Get(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		log.Println(err)
		return
	}
	post, err := postService.Get(&paramIdInt)
	if err != nil {
		json.NewEncoder(w).Encode(utils.ResponseError("Error getting the posts"))
		return
	}

	if post == nil || post.ID == 0 {
		json.NewEncoder(w).Encode(utils.ResponseError("Not found"))
	} else {
		json.NewEncoder(w).Encode(utils.ResponseSuccess("", post))
	}

}

// Store store
func (h *controller) Store(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		return
	}
	post.Created = time.Now()
	post.CreatedBy = 44000 // default for test

	resultId, err2 := postService.Create(&post)
	if err2 != nil {
		json.NewEncoder(w).Encode(utils.ResponseError("Error saving the post"))
		return
	}

	data := make(map[string]int)
	data["created_id"] = *resultId
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", data))
}

// Update update
func (h *controller) Update(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		log.Println(err)
	}
	var post models.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		return
	}
	result, err2 := postService.Update(&paramIdInt, &post)
	if err2 != nil {
		json.NewEncoder(w).Encode(utils.ResponseError("Error updating the post"))
		return
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess("", result))
}

// Remove remove
func (h *controller) Remove(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		log.Println(err)
	}
	resultId, err2 := postService.Remove(&paramIdInt)
	if err2 != nil {
		json.NewEncoder(w).Encode(utils.ResponseError("Error deleting the post"))
		return
	}
	data := make(map[string]int)
	data["deleted_id"] = *resultId
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", data))
}

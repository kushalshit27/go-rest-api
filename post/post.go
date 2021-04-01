package post

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/kushalshit27/go-rest-api/internal/database"
	"github.com/kushalshit27/go-rest-api/internal/models"
	"github.com/kushalshit27/go-rest-api/internal/utils"
)

type postAPI struct {
	db *database.DB
}

var (
	ctx = context.Background()
)

func newPostAPI(db *database.DB) *postAPI {
	var postapi *postAPI
	postapi = new(postAPI)
	postapi.db = db
	return postapi
}

// All all
func (h *postAPI) All(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query(ctx, "SELECT * FROM posts ORDER BY id DESC")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var results []models.Post
	for rows.Next() {
		var r models.Post
		err = rows.Scan(&r.ID, &r.Title, &r.Description, &r.Created, &r.Status)
		if err != nil {
			log.Println(err)
			return
		}
		results = append(results, r)
	}

	json.NewEncoder(w).Encode(utils.ResponseSuccess("", results))
}

// Get get
func (h *postAPI) Get(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		log.Println(err)
		return
	}
	var post models.Post
	query := `SELECT id,title,description,created_on, status FROM posts WHERE id =$1`
	err = h.db.QueryRow(ctx, query, paramIdInt).Scan(&post.ID, &post.Title, &post.Description, &post.Created, &post.Status)
	if err != nil {
		log.Println(err)
		return
	}
	if post.ID == 0 {
		json.NewEncoder(w).Encode(utils.ResponseError("Not found"))
	} else {
		json.NewEncoder(w).Encode(utils.ResponseSuccess("", post))
	}

}

// Store store
func (h *postAPI) Store(w http.ResponseWriter, r *http.Request) {
	var p models.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Println(err)
		return
	}
	p.Created = time.Now()

	sqlStatement := `INSERT INTO posts(title, description, created_on) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err = h.db.QueryRow(ctx, sqlStatement, p.Title, p.Description, p.Created).Scan(&id)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("New record ID is:", id)
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", "Created"))
}

// Update update
func (h *postAPI) Update(w http.ResponseWriter, r *http.Request) {
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
	sqlStatement := `UPDATE posts SET title=$1, description=$2, status=$3 WHERE id =$4 RETURNING id,created_on`
	err = h.db.QueryRow(ctx, sqlStatement, post.Title, post.Description, post.Status, paramIdInt).Scan(&post.ID, &post.Created)
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", post))
}

// Remove remove
func (h *postAPI) Remove(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	paramIdInt, err := strconv.Atoi(paramId)
	if err != nil {
		log.Println(err)
	}
	deletedId := 0
	sqlStatement := `DELETE FROM posts WHERE id =$1 RETURNING id`
	err = h.db.QueryRow(ctx, sqlStatement, paramIdInt).Scan(&deletedId)
	if err != nil {
		log.Println(err)
		return
	}
	data := make(map[string]int)
	data["deleted_id"] = deletedId
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", data))
}

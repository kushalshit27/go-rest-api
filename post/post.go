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
	postapi := new(postAPI)
	postapi.db = db
	return postapi
}

// All all
func (h *postAPI) All(w http.ResponseWriter, r *http.Request) {
	queryString := `SELECT 
						b.id,
						b.title,
						b.description,
						b.created_at, 
						b.status,
						u.name,
						u.email,
						u.role,
						u.created_at
					FROM 
						blogs as b 
					JOIN 
						users as u 
					ON 
					b.created_by = u.id;`

	rows, err := h.db.Query(ctx, queryString)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var results []models.Post
	for rows.Next() {
		var p models.Post
		err = rows.Scan(&p.ID, &p.Title, &p.Description, &p.Created, &p.Status, &p.User.Name, &p.User.Email, &p.User.Role, &p.User.CreatedAt)
		if err != nil {
			log.Println(err)
			return
		}
		results = append(results, p)
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
	query := `SELECT id,title,description,created_at, status FROM blogs WHERE id =$1`
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

	sqlStatement := `INSERT INTO blogs(title, description, created_on) VALUES ($1, $2, $3) RETURNING id`
	createdId := 0
	err = h.db.QueryRow(ctx, sqlStatement, p.Title, p.Description, p.Created).Scan(&createdId)
	if err != nil {
		log.Println(err)
		return
	}

	data := make(map[string]int)
	data["created_id"] = createdId
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", data))
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
	sqlStatement := `UPDATE blogs SET title=$1, description=$2, status=$3 WHERE id =$4 RETURNING id,created_on`
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
	sqlStatement := `DELETE FROM blogs WHERE id =$1 RETURNING id`
	err = h.db.QueryRow(ctx, sqlStatement, paramIdInt).Scan(&deletedId)
	if err != nil {
		log.Println(err)
		return
	}
	data := make(map[string]int)
	data["deleted_id"] = deletedId
	json.NewEncoder(w).Encode(utils.ResponseSuccess("", data))
}

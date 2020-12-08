package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post // storage
)

func init() {
	posts = []Post{Post{1, "Title 1", "text 1"}}
}

func GetPosts(res http.ResponseWriter, req *http.Request) {
	result, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the post array"}`))
		return
	}
	res.Header().Set("Content-type", "Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	var post Post

	for _, v := range posts {
		postID := strconv.Itoa(v.ID)
		if postID == id {
			//log.Println("Found")
			post = v
		}
	}
	result, _ := json.Marshal(post)
	res.Header().Set("Content-type", "Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}
func CreatePost(res http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the create post request"}`))
		return
	}
	post.ID = len(posts) + 1
	posts = append(posts, post)
	result, err := json.Marshal(posts)
	res.Header().Set("Content-type", "Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

/* TODO */
func UpdatePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"status": "TODO"}`))
	return
}

/* TODO */
func DeletePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "Application/json")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{"status": "TODO"}`))
	return
}

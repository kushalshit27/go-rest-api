package post

import (
	"github.com/kushalshit27/go-rest-api/internal/middleware"
	"github.com/kushalshit27/go-rest-api/internal/utils"

	"github.com/gorilla/mux"
)

// Service service
type Service struct {
	post *postAPI
}

// UserAPI
func PostAPI() *Service {
	return &Service{
		post: newPostAPI(),
	}
}

// Register Register
func (s *Service) Register(router *mux.Router) {
	router.Use(middleware.Logger)

	routes := utils.Routes{

		// Post
		utils.AddRoute("/posts", "GET", s.post.All),
		utils.AddRoute("/posts", "POST", s.post.Store),
		utils.AddRoute("/posts/{id}", "GET", s.post.Get),
		utils.AddRoute("/posts/{id}", "PUT", s.post.Update),
		utils.AddRoute("/posts/{id}", "DELETE", s.post.Remove),
	}

	for _, r := range routes {
		router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}
}
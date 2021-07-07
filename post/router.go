package post

import (
	"github.com/kushalshit27/go-rest-api/internal/database"
	"github.com/kushalshit27/go-rest-api/internal/middleware"
	"github.com/kushalshit27/go-rest-api/internal/utils"

	"github.com/gorilla/mux"
)

// Service service
type Service struct {
	db *database.DB
}

// UserAPI
func PostAPI(db *database.DB) *Service {
	return &Service{db}
}

// Register Register
func (s *Service) Register(router *mux.Router) {
	router.Use(middleware.Logger)

	var postRepository PostRepository = NewPostRepository(s.db)
	var postService PostService = NewPostService(postRepository)
	var postController PostController = NewPostController(postService)

	routes := utils.Routes{

		// Post
		utils.AddRoute("/posts", "GET", postController.All),
		utils.AddRoute("/posts", "POST", postController.Store),
		utils.AddRoute("/posts/{id}", "GET", postController.Get),
		utils.AddRoute("/posts/{id}", "PUT", postController.Update),
		utils.AddRoute("/posts/{id}", "DELETE", postController.Remove),
	}

	for _, r := range routes {
		router.HandleFunc(r.Path, r.Handler).Methods(r.Method)
	}
}

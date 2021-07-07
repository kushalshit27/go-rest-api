package post

import (
	"errors"

	"github.com/kushalshit27/go-rest-api/internal/models"
)

type PostService interface {
	Validate(post *models.Post) error
	FindAll() ([]models.Post, error)
	Create(post *models.Post) (*int, error)
	Get(id *int) (*models.Post, error)
	Update(id *int, post *models.Post) (*models.Post, error)
	Remove(id *int) (*int, error)
}

type service struct{}

var (
	repo PostRepository
)

func NewPostService(repository PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *models.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	return nil
}

func (*service) FindAll() ([]models.Post, error) {
	return repo.FindAll()
}

func (*service) Create(post *models.Post) (*int, error) {
	return repo.Save(post)
}

func (*service) Get(id *int) (*models.Post, error) {
	return repo.Get(id)
}

func (*service) Update(id *int, post *models.Post) (*models.Post, error) {
	return repo.Update(id, post)
}
func (*service) Remove(id *int) (*int, error) {
	return repo.Remove(id)
}

package post

import (
	"testing"

	"github.com/kushalshit27/go-rest-api/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *models.Post) (*int, error) {
	args := mock.Called()
	result := 1
	return &result, args.Error(1)
}

func (mock *MockRepository) FindAll() ([]models.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]models.Post), args.Error(1)
}

func (mock *MockRepository) Get(id *int) (*models.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*models.Post), args.Error(1)
}

func (mock *MockRepository) Update(id *int, post *models.Post) (*models.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*models.Post), args.Error(1)
}

func (mock *MockRepository) Remove(id *int) (*int, error) {
	args := mock.Called(id)
	result := 1
	return &result, args.Error(1)
}

func Test_FindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1

	post := models.Post{ID: identifier, Title: "A"}
	// Setup expectations
	mockRepo.On("FindAll").Return([]models.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	//Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
}

func Test_Create(t *testing.T) {
	mockRepo := new(MockRepository)
	post := models.Post{Title: "A"}

	var identifier int = 1

	//Setup expectations
	mockRepo.On("Save").Return(&identifier, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, identifier, *result)
	assert.Nil(t, err)
}

func Test_Validate_Empty_Post(t *testing.T) {

	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())

}

func Test_Validate_Empty_Post_Title(t *testing.T) {
	post := models.Post{ID: 1, Title: ""}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}

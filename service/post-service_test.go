package service

import (
	"testing"

	"com.github/fabiosebastiano/go-rest-api/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockRepository struct {
	mock.Mock
}

var (
	identifier int64 = 1
)

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	// setup exptecations
	post := entity.Post{
		ID:    identifier,
		Title: "A",
		Text:  "TEST",
	}
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.FindAll()

	require.NoError(t, err)

	//MOCK ASSERTION: BEHAVIOURAL
	mockRepo.AssertExpectations(t)

	require.NotEmpty(t, result)

	//Data assertion
	require.Equal(t, identifier, result[0].ID)
	require.Equal(t, "A", result[0].Title)
	require.Equal(t, "TEST", result[0].Text)

}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{
		ID:    identifier,
		Title: "A title",
		Text:  "Some random text",
	}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	require.NoError(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A title", result.Title)
	assert.Equal(t, "Some random text", result.Text)

}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.Error(t, err)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post object is empty")
}

func TestValidateEmptyTitle(t *testing.T) {
	testService := NewPostService(nil)

	post := entity.Post{
		ID:    1,
		Title: "",
		Text:  "TEST",
	}

	err := testService.Validate(&post)

	assert.Error(t, err)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post title is empty")
}

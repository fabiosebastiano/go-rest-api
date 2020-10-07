package service

import (
	"errors"
	"math/rand"
	"strconv"

	"com.github/fabiosebastiano/go-rest-api/entity"
	"com.github/fabiosebastiano/go-rest-api/repository"
)

// PostService interface
type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(postID string) (*entity.Post, error)
	Delete(postID string) error
}

type service struct{}

var (
	repo repository.PostRepository
)

// NewPostService inizializza l'interfaccia e restituisce il servizio
func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post object is empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)

}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (*service) FindByID(id string) (*entity.Post, error) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return repo.FindByID(id)
}

func (*service) Delete(id string) error {

	err := repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

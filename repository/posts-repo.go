package repository

import (
	"com.github/fabiosebastiano/go-rest-api/entity"
)

// PostRepository interface
type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(id string) (*entity.Post, error)
	Delete(post *entity.Post) error
}

package controller

import (
	"encoding/json"
	"net/http"

	"com.github/fabiosebastiano/go-rest-api/entity"
	"com.github/fabiosebastiano/go-rest-api/errors"
	"com.github/fabiosebastiano/go-rest-api/service"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error retrieving the posts array"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the post to save"})
		return
	}

	//VALIDAZIONE
	validationErr := postService.Validate(&post)
	if validationErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: validationErr.Error()})
		return
	}

	//CREAZIONE
	result, creationError := postService.Create(&post)
	if creationError != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error creating the post to save"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)

}

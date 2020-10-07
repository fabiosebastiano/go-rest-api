package controller

import (
	"encoding/json"
	"net/http"
	"strings"

	"com.github/fabiosebastiano/go-rest-api/cache"
	"com.github/fabiosebastiano/go-rest-api/entity"
	"com.github/fabiosebastiano/go-rest-api/errors"
	"com.github/fabiosebastiano/go-rest-api/service"
)

type controller struct{}

var (
	postService service.PostService
	postCache   cache.PostCache
)

type PostController interface {
	GetPostById(resp http.ResponseWriter, req *http.Request)
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService, cache cache.PostCache) PostController {
	postService = service
	postCache = cache
	return &controller{}
}

func (*controller) GetPostById(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	postID := strings.Split(req.URL.Path, "/")[2]

	var post *entity.Post = postCache.Get(postID)

	if post == nil {
		post, err := postService.FindByID(postID)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(errors.ServiceError{Message: "No post found with id" + postID})
			return
		}
		postCache.Set(postID, post)
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(post)
	} else {
		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(post)
	}

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

package controller

import (
	"bytes"
	json "encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"com.github/fabiosebastiano/go-rest-api/cache"
	"com.github/fabiosebastiano/go-rest-api/entity"
	"com.github/fabiosebastiano/go-rest-api/repository"
	"com.github/fabiosebastiano/go-rest-api/service"

	"github.com/stretchr/testify/assert"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postCacheSrv   cache.PostCache           = cache.NewRedisCache("localhost:6379", 0, 10)
	postController PostController            = NewPostController(postSrv, postCacheSrv)
)

const (
	ID    int64  = 1234
	TITLE string = "Just a title"
	TEXT  string = "Some random text here"
)

func TestAddPost(t *testing.T) {
	// 1) creare a ew HTTP POST request
	var jsonObject = []byte(`{"Title": "Just a title","Text": "Some random text here"}`)
	request, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonObject))

	// 2) assegnare func all'handler
	handler := http.HandlerFunc(postController.AddPost)

	// Record HTTP Response e httptest
	resp := httptest.NewRecorder()

	// Dispatch HTTP request
	handler.ServeHTTP(resp, request)

	// Add assertions on the status CODE e response
	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler ha risposto con codice non ok: ottenuto %v atteso %v", status, http.StatusOK)
	}

	var post entity.Post
	json.NewDecoder(resp.Body).Decode(&post)

	assert.NotNil(t, post.ID)
	assert.NotNil(t, post.Title)
	assert.NotNil(t, post.Text)

	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	//cleanup db
	tearDown(&post)
}

func TestGetPosts(t *testing.T) {

	setUp()

	request, _ := http.NewRequest("GET", "/posts", nil)

	// 2) assegnare func all'handler
	handler := http.HandlerFunc(postController.GetPosts)

	// Record HTTP Response e httptest
	resp := httptest.NewRecorder()

	// Dispatch HTTP request
	handler.ServeHTTP(resp, request)

	// Add assertions on the status CODE e response
	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler ha risposto con codice non ok: ottenuto %v atteso %v", status, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(resp.Body).Decode(&posts)

	assert.NotNil(t, posts[0].ID)
	assert.NotNil(t, posts[0].Title)
	assert.NotNil(t, posts[0].Text)

	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	//cleanup db
	tearDown(&posts[0])

}

func TestGetPostByID(t *testing.T) {

	setUp()

	request, _ := http.NewRequest("GET", "/posts/"+strconv.FormatInt(ID, 10), nil)

	// 2) assegnare func all'handler
	handler := http.HandlerFunc(postController.GetPostById)

	// Record HTTP Response e httptest
	resp := httptest.NewRecorder()

	// Dispatch HTTP request
	handler.ServeHTTP(resp, request)

	// Add assertions on the status CODE e response
	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler ha risposto con codice non ok: ottenuto %v atteso %v", status, http.StatusOK)
	}

	var post entity.Post
	json.NewDecoder(io.Reader(resp.Body)).Decode(&post)

	assert.NotNil(t, post.ID)
	assert.NotNil(t, post.Title)
	assert.NotNil(t, post.Text)

	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	//clean db
	tearDown(&post)

}

func tearDown(post *entity.Post) {
	postRepo.Delete(post)
}

func setUp() {
	var newPost entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepo.Save(&newPost)
}

package main

import (
	"fmt"
	"os"

	controller "com.github/fabiosebastiano/go-rest-api/controller"
	router "com.github/fabiosebastiano/go-rest-api/http"
	repository "com.github/fabiosebastiano/go-rest-api/repository"
	service "com.github/fabiosebastiano/go-rest-api/service"
)

var (
	//httpRouter     router.Router             = router.NewMuxRouter()
	httpRouter router.Router = router.NewMuxRouter()
	//postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postRepository repository.PostRepository = repository.NewSQLiteRepository()

	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	fmt.Println("PORTA RECUPERATA DA ENV: ", os.Getenv("PORT"))
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(os.Getenv("PORT"))

}

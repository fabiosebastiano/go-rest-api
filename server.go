package main

import (
	"fmt"
	"net/http"

	controller "com.github/fabiosebastiano/go-rest-api/controller"
	router "com.github/fabiosebastiano/go-rest-api/http"
	repository "com.github/fabiosebastiano/go-rest-api/repository"
	service "com.github/fabiosebastiano/go-rest-api/service"
)

var (
	//httpRouter     router.Router             = router.NewMuxRouter()
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, " UP & RUNNING")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)

}

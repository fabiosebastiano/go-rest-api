package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chiDispatcher = chi.NewRouter()
)

type chiRouter struct{}

//NewChiRouter .
func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, f)
}
func (*chiRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, f)
}
func (*chiRouter) SERVE(port string) {
	log.Printf("CHI Server listening on port %v", port)

	log.Fatalln(http.ListenAndServe(port, chiDispatcher))
}

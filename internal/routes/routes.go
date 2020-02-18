package routes

import (
	"log"
	"net/http"

	"github.com/centric-lt/restful-go/internal/controllers"
	"github.com/go-chi/chi"
)

func HandleRequests() {
	r := chi.NewRouter()
	r.Get("/", controllers.Welcome)
	r.Post("/articles", controllers.Posting)
	r.Get("/articles", controllers.AllArticles)
	log.Fatal(http.ListenAndServe(":8081", r))
}

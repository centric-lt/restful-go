package routes

import (
	"log"
	"net/http"

	"github.com/centric-lt/restful-go/internal/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", controllers.Welcome)
	http.HandleFunc("/posting", controllers.Posting)
	http.HandleFunc("/articles", controllers.AllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/centric-lt/restful-go/internal/article"
	"github.com/centric-lt/restful-go/internal/restdb"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "web/html/index.html")
	default:
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
}

func posting(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		postArticle := article.Article{Title: r.FormValue("title"), Description: r.FormValue("description"), Content: r.FormValue("content")}
		restdb.InsertArticle(postArticle)
		http.Redirect(w, r, "/", 301)
	default:
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := restdb.SelectAllArticles()
	fmt.Println("Endpoint: All Articles")
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/posting", posting)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	restdb.Connect()
	handleRequests()
	defer restdb.Close()
}

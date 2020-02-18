package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/centric-lt/restful-go/internal/models/article"
	"github.com/centric-lt/restful-go/internal/restdb"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/html/index.html")
}

func Posting(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	postArticle := article.Article{Title: r.FormValue("title"), Description: r.FormValue("description"), Content: r.FormValue("content")}
	restdb.InsertArticle(postArticle)
	http.Redirect(w, r, "/", 301)
}

func AllArticles(w http.ResponseWriter, r *http.Request) {
	articles := restdb.SelectAllArticles()
	fmt.Println("Endpoint: All Articles")
	json.NewEncoder(w).Encode(articles)
}

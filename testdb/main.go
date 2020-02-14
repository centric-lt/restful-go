package testdb

import (
	"database/sql"
	"fmt"
	article "restfulGo/article"

	_ "github.com/go-sql-driver/mysql"
)

var conn = dbConnection{driver: "mysql", credentials: "root:root@tcp(127.0.0.1:3306)/test"}

var db sql.DB

type dbConnection struct {
	driver      string
	credentials string
}

// Connect to database
func Connect() {
	cdb, err := sql.Open(conn.driver, conn.credentials)
	if err != nil {
		panic(err.Error())
	}
	db = *cdb
}

// Insert to database
func Insert(query string) {
	insert, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}

// InsertArticle to database
func InsertArticle(articleData article.Article) {
	fmt.Println(articleData)
	_, err := db.Exec("INSERT INTO articles (title, descr, content) VALUES (?, ?, ?)", articleData.Title, articleData.Description, articleData.Content)
	if err != nil {
		panic(err.Error())
	}
}

// Select from database
func Select(query string) *sql.Rows {
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	return results
}

// SelectAllArticles from database
func SelectAllArticles() article.Articles {
	results, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}
	var articlesData article.Articles
	for results.Next() {
		var articleData article.Article
		results.Scan(&articleData.Title, &articleData.Description, &articleData.Content)
		articlesData = append(articlesData, articleData)
	}
	//err = results.Scan(&article.Title, &article.Desc, &article.Content)
	if err != nil {
		panic(err.Error())
	}
	return articlesData
}

// Close connection to database
func Close() {
	db.Close()
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {

	// Connect
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Insert
	insert, err := db.Query("INSERT INTO articles VALUES ('test', 'testas is remote', 'irasau')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	// Select
	results, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var article Article
		err = results.Scan(&article.Title, &article.Desc, &article.Content)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(article.Desc)
	}

	fmt.Println("Success")
}

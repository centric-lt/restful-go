package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/derp?multiStatements=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations/sql/",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	m.Down()
	m.Up()

	fmt.Println("Migration complete")
}

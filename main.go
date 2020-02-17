package main

import (
	"github.com/centric-lt/restful-go/internal/restdb"
	"github.com/centric-lt/restful-go/internal/routes"
)

func main() {
	restdb.Connect()
	routes.HandleRequests()
	defer restdb.Close()
}

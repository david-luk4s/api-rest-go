package main

import (
	"fmt"

	"github.com/david-luk4s/go-api-rest/database"
	"github.com/david-luk4s/go-api-rest/routes"
)

func main() {
	database.ConnectionDB()

	fmt.Println("Iniciando servidor Rest com go")
	routes.HandleRouter()
}

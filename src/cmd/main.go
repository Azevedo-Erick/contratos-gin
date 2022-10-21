package main

import (
	"test/src/internal/database"
	"test/src/internal/routes"
)

func main() {

	database.ConnectDB()
	routes.Run()
}

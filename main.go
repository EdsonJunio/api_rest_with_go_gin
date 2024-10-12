package main

import (
	"github.com/EdsonJunio/api_rest_with_go_gin/database"
	"github.com/EdsonJunio/api_rest_with_go_gin/routes"
)

func main() {
	database.ConectaCoimBancoDeDados()
	routes.HandleResquests()
}

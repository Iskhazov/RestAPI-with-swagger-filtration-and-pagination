// @title Test Task API for Effective Mobile
// @version 1.0
// @description REST API for managing a list of people with support for filtering, pagination, and CRUD operations.
// @host localhost:8080
// @BasePath /api/v1
package main

import (
	"awesomeProject2/cmd/api"
	"awesomeProject2/config"
	"awesomeProject2/storage"
	"fmt"
	"log"
)

func main() {
	db, err := storage.MyNewSQlStorage(config.Config{
		DBHost: config.Envs.DBHost,
		DBUser: config.Envs.DBUser,
		DBPass: config.Envs.DBPass,
		DBPort: config.Envs.DBPort,
		DBName: config.Envs.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to database %s\n", config.Envs.DBName)
	//start server
	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

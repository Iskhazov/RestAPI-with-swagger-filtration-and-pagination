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
		PublicHost: config.Envs.PublicHost,
		DBUser:     config.Envs.DBUser,
		DBPass:     config.Envs.DBPass,
		Port:       config.Envs.Port,
		DBName:     config.Envs.DBName,
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

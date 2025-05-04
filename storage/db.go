package storage

import (
	"awesomeProject2/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func MyNewSQlStorage(cfg config.Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName))
	if err != nil {
		log.Printf("error connecting to DB: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

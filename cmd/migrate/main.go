package main

import (
	"awesomeProject2/config"
	"awesomeProject2/storage"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"os"
)

func main() {
	// Initialize database connection using environment variables
	db, err := storage.MyNewSQlStorage(config.Config{
		DBHost: config.Envs.DBHost,
		DBUser: config.Envs.DBUser,
		DBPass: config.Envs.DBPass,
		DBPort: config.Envs.DBPort,
		DBName: config.Envs.DBName,
	})
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect to the database: %v", err)
	}

	// Create a new migration driver instance for Postgres
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("[ERROR] Failed to create migration driver: %v", err)
	}

	// Initialize the migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // path to migration files
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("[ERROR] Failed to initialize migrations: %v", err)
	}

	// Retrieve the migration command (last CLI argument)
	cmd := os.Args[(len(os.Args) - 1)]

	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("[ERROR] Migration up failed: %v", err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("[ERROR] Migration down failed: %v", err)
		}
	}
}

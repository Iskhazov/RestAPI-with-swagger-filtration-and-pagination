package api

import (
	_ "awesomeProject2/docs" // Import generated Swagger docs
	"awesomeProject2/service"
	"awesomeProject2/storage"
	"database/sql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

type Server struct {
	addr string
	db   *sql.DB
}

// NewServer initializes a new HTTP server with given address and DB connection
func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

// Run starts the HTTP server and sets up routing
func (s *Server) Run() error {
	log.Println("[INFO] Starting server at", s.addr)

	// Initialize main router
	router := mux.NewRouter()

	// Attach Swagger UI handler (available at /swagger/index.html)
	log.Println("[DEBUG] Registering Swagger handler at /swagger/")
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Create subrouter for API versioning
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize storage layer
	personStore := storage.NewStore(s.db)
	log.Println("[DEBUG] Initialized person storage layer")

	// Initialize service layer with storage
	personService := service.NewLayerService(personStore)
	log.Println("[DEBUG] Initialized person service layer")

	// Initialize request handler with service
	requestHandler := service.NewHandler(personService)
	log.Println("[DEBUG] Registering HTTP handlers for person service")

	// Register HTTP routes
	requestHandler.Routes(subrouter)

	log.Println("[INFO] Server is listening on", s.addr)

	// Start HTTP server
	err := http.ListenAndServe(s.addr, router)
	if err != nil {
		log.Printf("[ERROR] Failed to start server: %v\n", err)
	}
	return err
}

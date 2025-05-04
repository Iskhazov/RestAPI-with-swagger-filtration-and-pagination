package api

import (
	_ "awesomeProject2/docs"
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

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

func (s *Server) Run() error {
	log.Println("Starting server at ", s.addr)
	router := mux.NewRouter()

	// Swagger handler
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	personStore := storage.NewStore(s.db)
	personService := service.NewLayerService(personStore)
	requestHandler := service.NewHandler(personService)

	requestHandler.Routes(subrouter)

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}

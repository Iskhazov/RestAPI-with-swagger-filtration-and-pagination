package api

import (
	"awesomeProject2/service"
	"awesomeProject2/storage"
	"database/sql"
	"github.com/gorilla/mux"
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
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	personStore := storage.NewStore(s.db)
	personService := service.NewLayerService(personStore)
	requestHandler := service.NewHandler(personService)

	requestHandler.Routes(subrouter)

	log.Println("Listening on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}

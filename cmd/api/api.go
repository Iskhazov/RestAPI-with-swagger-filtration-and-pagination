package api

import (
	"database/sql"
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
	return http.ListenAndServe(s.addr, nil)
}

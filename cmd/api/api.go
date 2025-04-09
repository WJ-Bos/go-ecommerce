package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	address string
	db      *sql.DB
}

// NewAPIServer a New Instance of the API Server
func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

// Run the Api Server
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log.Println("API v1 server listening on " + s.address)

	return http.ListenAndServe(s.address, router)
}

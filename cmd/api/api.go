package api

import (
	"Ecom/service/user"
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
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	//creating user service handler
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	//passing the subRouter to Register the Routes with the Correct prefix
	userHandler.RegisterRoutes(subRouter)

	log.Println("API v1 server listening on " + s.address)

	return http.ListenAndServe(s.address, router)
}

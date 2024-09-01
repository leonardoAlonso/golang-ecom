// By convention in go we create a folder called cmd to the entry point of the application
package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/leonardoAlonso/go-ecom/services/user"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	/*
	   This return a pointer to the struct because
	   we want to modify the struct in the future
	   in the same memory address
	*/
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	/*
	   This initialize the server and start the server
	   with the address and the router
	   uses a pointer to the struct because we want to operate
	   on the same memory address
	*/
	// Start the server
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new handler
	userHandler := user.NewHandler()
	// Register the routes
	userHandler.RegisterRoutes(subrouter)

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

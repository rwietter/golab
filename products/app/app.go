package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// CreateServer creates a new server
func CreateServer() *http.Server {
	r := mux.NewRouter()

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return server
}

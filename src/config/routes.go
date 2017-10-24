package config

import (
	"handlers/api/v1"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	println("Initializing routes")

	r := mux.NewRouter()

	r.HandleFunc("/api/v1", v1.HomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users", v1.UsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id:[0-9]+}", v1.UserHandler).Methods(http.MethodGet)

	http.Handle("/", r)

	println("Routes initialized")
}

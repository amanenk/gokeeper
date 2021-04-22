package handlers

import (
	"github.com/fdistorted/gokeeper/handlers/middlewares"
	tables2 "github.com/fdistorted/gokeeper/handlers/tables"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false) //todo move it to config

	r.Use(middlewares.RequestID)

	tables := r.PathPrefix("/tables").Subrouter()
	tables.HandleFunc("/", tables2.GetAll).Methods(http.MethodGet)
	tables.HandleFunc("/{id}", tables2.Get).Methods(http.MethodGet)
	tables.HandleFunc("/{id}", tables2.Put).Methods(http.MethodPut)

	return r
}

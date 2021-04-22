package handlers

import (
	"fmt"
	"github.com/fdistorted/gokeeper/handlers/middlewares"
	"github.com/fdistorted/gokeeper/handlers/middlewares/role"
	tables2 "github.com/fdistorted/gokeeper/handlers/tables"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false) //todo move it to config

	r.Use(middlewares.RequestID)

	tables := r.PathPrefix("/tables").Subrouter()
	tables.HandleFunc("/", tables2.GetAll).Methods(http.MethodGet)
	tables.HandleFunc("/{id}", tables2.Get).Methods(http.MethodGet)
	tables.HandleFunc("/{id}", tables2.Put).Methods(http.MethodPut)

	//will be used to make meals ready
	admin := r.PathPrefix("/admin").Subrouter()
	admin.Use(role.NewRoleFilter(role.Admin).Attach)
	admin.Use(middlewares.JWT)
	admin.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "admin api %d\n", time.Now().Unix())
	}).Methods(http.MethodGet)

	return r
}

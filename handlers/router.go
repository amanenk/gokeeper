package handlers

import (
	"fmt"
	"github.com/fdistorted/gokeeper/handlers/middlewares"
	"github.com/fdistorted/gokeeper/handlers/middlewares/role"
	"github.com/fdistorted/gokeeper/handlers/orders"
	"github.com/fdistorted/gokeeper/handlers/tables"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(middlewares.RequestID)

	tablesRouter := r.PathPrefix("/tables").Subrouter()
	tablesRouter.HandleFunc("/", tables.GetAll).Methods(http.MethodGet)
	tablesRouter.HandleFunc("/{id}", tables.Get).Methods(http.MethodGet)
	tablesRouter.HandleFunc("/{id}", tables.Put).Methods(http.MethodPut)

	ordersRouter := r.PathPrefix("/orders").Subrouter()
	ordersRouter.HandleFunc("/", orders.GetAll).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/", orders.Post).Methods(http.MethodPost)
	ordersRouter.HandleFunc("/", orders.GetAll).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/", orders.GetAll).Methods(http.MethodGet)

	//will be used to trigger meals as ready
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(role.NewRoleFilter(role.Admin).Attach)
	adminRouter.Use(middlewares.JWT)
	adminRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "admin api %d\n", time.Now().Unix())
	}).Methods(http.MethodGet)

	return r
}

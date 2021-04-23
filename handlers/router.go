package handlers

import (
	"fmt"
	"github.com/fdistorted/gokeeper/handlers/guests"
	"github.com/fdistorted/gokeeper/handlers/login"
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

	r.HandleFunc("/login", login.Post).Methods(http.MethodPost)

	waiterRoleFilter := role.NewRoleFilter(role.Waiter)
	tablesRouter := r.PathPrefix("/tables").Subrouter()
	tablesRouter.Use(waiterRoleFilter.Attach)
	tablesRouter.Use(middlewares.JWT)
	tablesRouter.HandleFunc("/", tables.GetAll).Methods(http.MethodGet)
	tablesRouter.HandleFunc("/{id}", tables.Get).Methods(http.MethodGet)
	tablesRouter.HandleFunc("/{id}", tables.Put).Methods(http.MethodPut)

	ordersRouter := r.PathPrefix("/orders").Subrouter()
	ordersRouter.Use(waiterRoleFilter.Attach)
	ordersRouter.Use(middlewares.JWT)
	ordersRouter.HandleFunc("/", orders.GetAll).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/", orders.Post).Methods(http.MethodPost)

	//host/orders/{orderid}/guests
	guestsRouter := ordersRouter.PathPrefix("/{orderId}/guests").Subrouter()
	guestsRouter.HandleFunc("/", guests.Post).Methods(http.MethodPost)
	guestsRouter.HandleFunc("/{guestId}", guests.Delete).Methods(http.MethodDelete)

	//will be used to trigger meals as ready
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(role.NewRoleFilter(role.Admin).Attach)
	adminRouter.Use(middlewares.JWT)
	adminRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "admin api %d\n", time.Now().Unix())
	}).Methods(http.MethodGet)

	return r
}

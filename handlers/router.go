package handlers

import (
	"fmt"
	bill_items "github.com/fdistorted/gokeeper/handlers/bill-items"
	"github.com/fdistorted/gokeeper/handlers/bills"
	"github.com/fdistorted/gokeeper/handlers/guests"
	"github.com/fdistorted/gokeeper/handlers/login"
	"github.com/fdistorted/gokeeper/handlers/meals"
	"github.com/fdistorted/gokeeper/handlers/middlewares"
	"github.com/fdistorted/gokeeper/handlers/middlewares/role"
	order_Items "github.com/fdistorted/gokeeper/handlers/order-Items"
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
	mealsRouter := r.PathPrefix("/meals").Subrouter()
	mealsRouter.Use(waiterRoleFilter.Attach)
	mealsRouter.Use(middlewares.JWT)
	mealsRouter.HandleFunc("/", meals.GetAll).Methods(http.MethodGet)

	tablesRouter := r.PathPrefix("/tables").Subrouter()
	tablesRouter.Use(waiterRoleFilter.Attach)
	tablesRouter.Use(middlewares.JWT)
	tablesRouter.HandleFunc("/", tables.GetAll).Methods(http.MethodGet)
	tablesRouter.HandleFunc("/{tableId}", tables.Get).Methods(http.MethodGet)
	tablesRouter.HandleFunc("/{tableId}", tables.Put).Methods(http.MethodPut)

	ordersRouter := r.PathPrefix("/orders").Subrouter()
	ordersRouter.Use(waiterRoleFilter.Attach)
	ordersRouter.Use(middlewares.JWT)
	ordersRouter.HandleFunc("/", orders.GetAll).Methods(http.MethodGet)
	ordersRouter.HandleFunc("/", orders.Post).Methods(http.MethodPost)
	ordersRouter.HandleFunc("/{orderId}/finish", orders.Finish).Methods(http.MethodPost)

	//host/orders/{orderid}/guests
	guestsRouter := ordersRouter.PathPrefix("/{orderId}/guests").Subrouter()
	guestsRouter.HandleFunc("/", guests.Post).Methods(http.MethodPost)
	guestsRouter.HandleFunc("/{guestId}", guests.Delete).Methods(http.MethodDelete)

	//host/orders/{orderid}/guests/{billId}/order-items to manage guest ordered items
	orderGuestItemsRouter := guestsRouter.PathPrefix("/{orderId}/guests/{guestId}").Subrouter()
	orderGuestItemsRouter.HandleFunc("/{orderedItemId}", guests.Post).Methods(http.MethodPost)
	orderGuestItemsRouter.HandleFunc("/{orderedItemId}", guests.Delete).Methods(http.MethodDelete)

	//host/orders/{orderid}/order-items
	orderItemsRouter := ordersRouter.PathPrefix("/{orderId}/order-items").Subrouter()
	orderItemsRouter.HandleFunc("/", order_Items.Post).Methods(http.MethodPost)
	orderItemsRouter.HandleFunc("/{orderedItemId}", guests.Delete).Methods(http.MethodDelete)

	//host/orders/{orderid}/bills
	orderBillsRouter := ordersRouter.PathPrefix("/{orderId}/bills").Subrouter()
	orderBillsRouter.HandleFunc("/", bills.Post).Methods(http.MethodPost)
	orderBillsRouter.HandleFunc("/{billId}", bills.Delete).Methods(http.MethodDelete)

	//host/orders/{orderId}/bills/{billId}/order-items to manage bill ordered items
	orderBillItemsRouter := orderBillsRouter.PathPrefix("/{orderId}/bills/{billId}").Subrouter()
	orderBillItemsRouter.HandleFunc("/{orderedItemId}", bill_items.Post).Methods(http.MethodPost)
	orderBillItemsRouter.HandleFunc("/{orderedItemId}", bill_items.Delete).Methods(http.MethodDelete)

	//will be used to mark meals as ready
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(role.NewRoleFilter(role.Admin).Attach)
	adminRouter.Use(middlewares.JWT)
	adminRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "admin api %d\n", time.Now().Unix())
	}).Methods(http.MethodGet)

	return r
}

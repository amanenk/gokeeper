package bill_items

import (
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//orderId, ok := vars["orderId"]
	//if !ok {
	//	logger.Get().Error("missing parameter")
	//	common.SendError(w, errorTypes.NewNoFieldError("id"))
	//	return
	//}

	//todo check if order is owned by a waiter

	//create the
	//todo check if order is not billed yet
	//todo unattach item from previous guest
	//todo attach item to a new guest

	common.SendError(w, errorTypes.NewNotImplemented())
}

package bill_items

import (
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//orderId, ok := vars["orderId"]
	//if !ok {
	//	logger.Get().Error("missing parameter")
	//	common.SendError(w, errorTypes.NewNoFieldError("orderId"))
	//	return
	//}
	//
	////todo check if order is owned by a waiter
	//
	//orderedItemId, ok := vars["orderedItemId"]
	//if !ok {
	//	logger.Get().Error("missing parameter")
	//	common.SendError(w, errorTypes.NewNoFieldError("orderedItemId"))
	//	return
	//}

	//todo check if order is not billed yet
	//todo unattach item from guest

	common.SendError(w, errorTypes.NewNotImplemented())
}

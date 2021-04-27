package bills

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/bill"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billId, ok := vars["billId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("orderId"))
		return
	}
	orderObj, err := common.GetOrderEditableByWaiter(w, r)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("problems getting users order", zap.Error(err))
	}

	var bill bill.Bill
	tx := database.Get().Preload("ordered_meals").First(bill, billId)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendResponse(w, http.StatusOK, orderObj.Bills[0])
	//todo remove all items from bill

	common.SendError(w, errorTypes.NewNotImplemented())
}

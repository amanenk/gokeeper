package orders

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/order"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Finish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["orderId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("id"))
		return
	}

	var order order.Order

	waiterId := 1 //todo retrieve userId from context
	chain := database.Get().Where("waiter_id = ?", waiterId)

	if err := chain.Find(&order, orderId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	//todo check orders if they are ready to finish the order and finish it

	common.SendResponse(w, http.StatusOK, order)
}

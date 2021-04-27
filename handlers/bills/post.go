package bills

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/logger"
	bill2 "github.com/fdistorted/gokeeper/models/bill"
	"go.uber.org/zap"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	orderObj, err := common.GetOrderEditableByWaiter(w, r)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("problems getting users order", zap.Error(err))
	}

	var bill bill2.Bill
	orderObj.Bills = append(orderObj.Bills, bill)
	tx := database.Get().Save(orderObj)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendResponse(w, http.StatusOK, orderObj.Bills[0])
}

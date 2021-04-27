package bills

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/bill"
	ordered_meal "github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	billId, ok := vars["billId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("billId"))
		return
	}
	orderObj, err := common.GetOrderEditableByWaiter(w, r)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("problems getting users order", zap.Error(err))
	}

	var billObj bill.Bill
	tx := database.Get().
		Preload("OrderedMeals").
		Where("order_id = ?", orderObj.ID).
		First(&billObj, billId)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	for _, v := range billObj.OrderedMeals {
		v.GuestID = nil
		tx = database.Get().
			Model(ordered_meal.OrderedMeal{}).
			Where("bill_id = ?", billObj.ID).
			Updates(map[string]interface{}{"bill_id": nil})
		if tx.Error != nil {
			tx.Rollback()
			common.HandleDatabaseError(w, tx.Error)
			return
		}
	}

	tx = tx.Delete(&billObj)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendOk(w)
}

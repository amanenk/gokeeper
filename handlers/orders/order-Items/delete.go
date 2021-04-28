package order_Items

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	ordered_meal "github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderObj, err := common.GetOrderEditableByWaiter(w, r)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("problems getting users order", zap.Error(err))
	}

	orderedItemId, ok := vars["orderedItemId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("orderedItemId"))
		return
	}

	var orderedItemObj ordered_meal.OrderedMeal
	tx := database.Get().WithContext(r.Context()).Find(&orderedItemObj, orderedItemId)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	if orderedItemObj.Status != ordered_meal.MealOrdered {
		common.SendError(w, errorTypes.NewBadRequestError("meal cant be removed already"))
	}

	tx = tx.Where("order_id = ?", orderObj.ID).
		Delete(&orderedItemObj)
	if tx.Error != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendOk(w)
}

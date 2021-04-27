package guests

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/guest"
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

	guestId, ok := vars["guestId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("guestId"))
		return
	}

	var guestObj guest.Guest
	tx := database.Get().
		Preload("OrderedMeals").
		Where("order_id = ?", orderObj.ID).
		Find(&guestObj, guestId)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	for _, v := range guestObj.OrderedMeals {
		v.GuestID = nil
		tx = database.Get().
			Model(ordered_meal.OrderedMeal{}).
			Where("guest_id = ?", guestObj.ID).
			Updates(map[string]interface{}{"guest_id": nil})
		if tx.Error != nil {
			tx.Rollback()
			common.HandleDatabaseError(w, tx.Error)
			return
		}
	}

	tx = tx.Delete(guestObj)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendOk(w)
}

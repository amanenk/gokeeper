package orders

import (
	"database/sql"
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/jwt"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/order"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Finish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	orderId, ok := vars["orderId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("id"))
		return
	}

	waiterId, err := jwt.GetUserId(r.Context())
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("failed to get waiterId")
		common.SendError(w, errorTypes.NewUnauthorized())
		return
	}

	var orderObj *order.Order
	if err := database.Get().
		Preload("Bills.OrderedMeals.Meal").
		Preload("OrderedMeals").
		Where("waiter_id = ?", waiterId).
		Find(&orderObj, orderId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	if orderObj.Status == order.StatusBilled || orderObj.Status == order.StatusFinished {
		common.SendError(w, errorTypes.NewBadRequestError("order is not changable already"))
		return
	}

	for _, item := range orderObj.OrderedMeals {
		if item.BillID == nil {
			common.SendError(w, errorTypes.NewBadRequestError("please add all the items to the bills"))
			return
		}
	}

	for i := range orderObj.Bills {
		if len(orderObj.Bills[i].OrderedMeals) == 0 {
			common.SendError(w, errorTypes.NewBadRequestError("there are empty bills remove them"))
			return
		}
		var total uint
		for _, item := range orderObj.Bills[i].OrderedMeals {
			total += item.Amount * item.Meal.Price
		}
		orderObj.Bills[i].Total = total
	}
	// return list of bills

	orderObj.Status = order.StatusBilled
	orderObj.FinishedAt = sql.NullTime{Time: time.Now(), Valid: true}

	tx := database.Get().Save(orderObj)
	if tx.Error != nil {
		tx.Rollback()
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendResponse(w, http.StatusOK, orderObj.Bills)
}

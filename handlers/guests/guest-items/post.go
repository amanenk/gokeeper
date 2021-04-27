package guest_items

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/jwt"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/guest"
	"github.com/fdistorted/gokeeper/models/order"
	ordered_meal "github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["orderId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("orderId"))
		return
	}

	waiterId, err := jwt.GetUserId(r.Context())
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("failed to get waiterId")
		common.SendError(w, errorTypes.NewUnauthorized())
	}

	var orderObj order.Order
	if err := database.Get().Where("waiter_id = ?", waiterId).First(&orderObj, orderId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	if orderObj.Status == order.StatusBilled || orderObj.Status == order.StatusFinished {
		common.SendError(w, errorTypes.NewBadRequestError("order is not changable already"))
	}

	orderedItemId, ok := vars["orderedItemId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("orderedItemId"))
		return
	}

	guestId, ok := vars["guestId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("guestId"))
		return
	}

	var billObj guest.Guest
	if err := database.Get().Where("order_id = ?", orderId).First(&billObj, guestId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	var orderedMealObj ordered_meal.OrderedMeal
	if err := database.Get().Where("guest_id = ?", guestId).First(&orderedMealObj, orderedItemId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	uint64GuestId, _ := strconv.ParseUint(guestId, 10, 32)
	uintGuestId := uint(uint64GuestId)
	orderedMealObj.GuestID = &uintGuestId

	if err := database.Get().Save(orderedMealObj).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendOk(w)
}

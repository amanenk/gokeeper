package common

import (
	"errors"
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/jwt"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/order"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

func GetOrderEditableByWaiter(w http.ResponseWriter, r *http.Request) (*order.Order, error) {
	vars := mux.Vars(r)
	orderId, ok := vars["orderId"]
	if !ok {
		logger.Get().Error("missing parameter")
		SendError(w, errorTypes.NewNoFieldError("id"))
		return nil, errors.New("no order id provided")
	}

	waiterId, err := jwt.GetUserId(r.Context())
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("failed to get waiterId")
		SendError(w, errorTypes.NewUnauthorized())
		return nil, errors.New("no waiterId in context")
	}

	//todo add some caching here
	var orderObj *order.Order
	if err := database.Get().Where("waiter_id = ?", waiterId).First(&orderObj, orderId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		HandleDatabaseError(w, err)
		return nil, errors.New("such a waiter does not have this orded")
	}

	if orderObj.Status == order.StatusBilled || orderObj.Status == order.StatusFinished {
		SendError(w, errorTypes.NewBadRequestError("order is not changable already"))
		return nil, errors.New("order is closed")
	}
	return orderObj, err
}

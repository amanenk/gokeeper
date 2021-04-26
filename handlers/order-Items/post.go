package order_Items

import (
	"errors"
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	ordered_meal "github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/fdistorted/gokeeper/validator"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["orderId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("id"))
		return
	}

	//todo check if order is owned by a waiter

	var orderItemObj ordered_meal.OrderedMeal

	if jsonError := common.UnmarshalRequestBody(r, &orderItemObj); jsonError != nil {
		common.SendError(w, *jsonError)
		return
	}

	if err := validator.Get().Struct(&orderItemObj); err != nil {
		logger.WithCtxValue(r.Context()).Error("the object did not pass the validation", zap.Error(err))
		common.HandleValidationError(w, err)
		return
	}

	orderIdNumber, err := strconv.ParseUint(orderId, 10, 32)
	if err != nil {
		common.SendError(w, errorTypes.NewBadRequestError("orderId is not a number"))
		return
	}

	if orderItemObj.Amount < 1 {
		common.SendError(w, errorTypes.NewBadRequestError("wrong amount of meals ordered"))
		return
	}

	orderItemObj.Status = ordered_meal.MealOrdered
	orderItemObj.OrderId = uint(orderIdNumber)

	tx := database.Get().Find(&orderItemObj.Meal, orderItemObj.Meal.ID)
	if tx.Error != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.SendError(w, errorTypes.NewBadRequestError("ordered meal does not exist"))
		}
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	tx = database.Get().Create(orderItemObj)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendResponse(w, http.StatusOK, orderItemObj)
}

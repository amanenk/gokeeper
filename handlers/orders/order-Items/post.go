package order_Items

import (
	"errors"
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	ordered_meal "github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/fdistorted/gokeeper/validator"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	orderObj, err := common.GetOrderEditableByWaiter(w, r)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("problems getting users order", zap.Error(err))
	}

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

	if orderItemObj.Amount < 1 {
		common.SendError(w, errorTypes.NewBadRequestError("wrong amount of meals ordered"))
		return
	}

	tx := database.Get().
		WithContext(r.Context()).
		First(&orderItemObj.Meal, orderItemObj.Meal.ID)
	if tx.Error != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.SendError(w, errorTypes.NewBadRequestError("ordered meal does not exist"))
			return
		}
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	orderItemObj.Status = ordered_meal.MealOrdered
	orderObj.OrderedMeals = append(orderObj.OrderedMeals, orderItemObj)

	tx = tx.Save(orderObj)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendResponse(w, http.StatusOK, orderObj.OrderedMeals[0])
}

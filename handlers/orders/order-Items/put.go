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

//todo implement or remove it
//used to update busy status of the table
func Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := common.GetOrderEditableByWaiter(w, r)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("problems getting users order", zap.Error(err))
	}

	orderedItemId, ok := vars["orderedItemId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("orderedItemId"))
		return
	}

	var orderedMealObj ordered_meal.OrderedMeal
	tx := database.Get().Find(&orderedMealObj, orderedItemId)
	if tx.Error != nil {
		tx.Rollback()
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	if jsonError := common.UnmarshalRequestBody(r, &orderedMealObj); jsonError != nil {
		tx.Rollback()
		common.SendError(w, *jsonError)
		return
	}

	//if err := validator.Get().Struct(&orderedMealObj); err != nil {
	//	tx.Rollback()
	//	logger.WithCtxValue(r.Context()).Error("data did not pass validation", zap.Error(err))
	//	common.HandleValidationError(w, err)
	//	return
	//}

	if err := database.Get().Save(&orderedMealObj).Error; err != nil {
		tx.Rollback()
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendResponse(w, http.StatusOK, orderedMealObj)
}

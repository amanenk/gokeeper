package orders

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/order"
	"github.com/fdistorted/gokeeper/validator"
	"go.uber.org/zap"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	//todo retrieve userId from context

	var order order.Order
	if jsonError := common.UnmarshalRequestBody(r, &order); jsonError != nil {
		common.SendError(w, *jsonError)
		return
	}

	if err := validator.Get().Struct(&order); err != nil {
		logger.WithCtxValue(r.Context()).Error("the object did not pass the validation", zap.Error(err))
		common.HandleValidationError(w, err)
		return
	}

	if err := database.Get().Create(&order).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendResponse(w, http.StatusOK, order)
}

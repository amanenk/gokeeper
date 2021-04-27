package orders

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/jwt"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/guest"
	"github.com/fdistorted/gokeeper/models/order"
	"github.com/fdistorted/gokeeper/validator"
	"go.uber.org/zap"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var orderObj order.Order
	if jsonError := common.UnmarshalRequestBody(r, &orderObj); jsonError != nil {
		common.SendError(w, *jsonError)
		return
	}

	waiterId, err := jwt.GetUserId(r.Context())
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("failed to get waiterId")
		common.SendError(w, errorTypes.NewUnauthorized())
	}

	//set initial values
	orderObj.WaiterID = waiterId
	orderObj.Status = order.StatusCreated
	orderObj.Guests = append(orderObj.Guests, guest.Guest{})

	if err := validator.Get().Struct(&orderObj); err != nil {
		logger.WithCtxValue(r.Context()).Error("the object did not pass the validation", zap.Error(err))
		common.HandleValidationError(w, err)
		return
	}

	tx := database.Get().Create(&orderObj)
	if tx.Error != nil {
		tx.Rollback()
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendResponse(w, http.StatusOK, orderObj)
}

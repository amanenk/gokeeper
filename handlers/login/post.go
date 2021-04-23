package login

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/handlers/middlewares/role"
	jwt3 "github.com/fdistorted/gokeeper/jwt"
	"github.com/fdistorted/gokeeper/logger"
	jwt2 "github.com/fdistorted/gokeeper/models/jwt"
	"github.com/fdistorted/gokeeper/models/waiter"
	"github.com/fdistorted/gokeeper/validator"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var waiterObj, retrievedWaiter waiter.Waiter
	if jsonError := common.UnmarshalRequestBody(r, &waiterObj); jsonError != nil {
		common.SendError(w, *jsonError)
		return
	}

	if err := validator.Get().Struct(&waiterObj); err != nil {
		logger.WithCtxValue(r.Context()).Error("the object did not pass the validation", zap.Error(err))
		common.HandleValidationError(w, err)
		return
	}

	tx := database.Get().Where("email = ?", waiterObj.Email).Find(&retrievedWaiter)
	if tx.Error != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(tx.Error))
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	if retrievedWaiter.Email == waiterObj.Email && retrievedWaiter.Password == waiterObj.Password {
		token, err := jwt3.NewToken(strconv.FormatUint(uint64(retrievedWaiter.ID), 10), string(role.Waiter))
		if err != nil {
			common.SendError(w, errorTypes.NewInternalServerError())
			return
		}
		// todo add refresh token
		common.SendResponse(w, http.StatusOK, jwt2.JWTResponse{Token: token})
		return
	}

	common.SendError(w, errorTypes.NewUnauthorized())
}

package orders

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/jwt"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/order"
	"go.uber.org/zap"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var orders []order.Order

	waiterId, err := jwt.GetUserId(r.Context())
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("failed to get waiterId")
		common.SendError(w, errorTypes.NewUnauthorized())
	}
	chain := database.Get().Where("waiter_id = ?", waiterId)

	status := r.URL.Query().Get("status")
	if status != "" { //TODO add validation to check if status is one of the available statuses before doing the query
		chain.Where("status = ?", status)
	}

	if err := chain.Preload("Bills").Preload("OrderedMeals").Preload("Guests").Find(&orders).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendResponse(w, http.StatusOK, orders)
}

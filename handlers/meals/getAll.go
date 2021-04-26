package meals

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/meal"
	"go.uber.org/zap"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	// todo add filters if needed
	var meals []meal.Meal

	if err := database.Get().Find(&meals).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendResponse(w, http.StatusOK, meals)
}

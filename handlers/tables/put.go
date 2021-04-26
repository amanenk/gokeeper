package tables

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/table"
	"github.com/gorilla/mux"
	"net/http"
)

//used to update busy status of the table
func Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["tableId"]
	if !ok {
		logger.WithCtxValue(r.Context()).Error("missing id in request")
		common.SendError(w, errorTypes.NewNoFieldError("id"))
		return
	}

	var table table.Table

	tx := database.Get().Find(&table, id)
	if tx.Error != nil {
		tx.Rollback()
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	if jsonError := common.UnmarshalRequestBody(r, &table); jsonError != nil {
		tx.Rollback()
		common.SendError(w, *jsonError)
		return
	}

	//if err := validator.Get().Struct(&table); err != nil {
	//	tx.Rollback()
	//	logger.WithCtxValue(r.Context()).Error("data did not pass validation", zap.Error(err))
	//	common.HandleValidationError(w, err)
	//	return
	//}

	if err := database.Get().Save(&table).Error; err != nil {
		tx.Rollback()
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendResponse(w, http.StatusOK, table)
}

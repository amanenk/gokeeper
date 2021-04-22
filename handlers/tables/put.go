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
	id, ok := vars["id"]
	if !ok {
		logger.WithCtxValue(r.Context()).Error("missing id in request")
		common.SendError(w, errorTypes.NewNoFieldError("id"))
		return
	}

	var table, dbTable table.Table
	jsonError := common.UnmarshalRequestBody(r, &table)
	if jsonError != nil {
		common.SendError(w, errorTypes.NewNoFieldError("id"))
		return
	}

	tx := database.DB.Model(&dbTable).Where("id = ?", id).Updates(map[string]interface{}{"is_busy": table.IsBusy})

	if tx.Error != nil {
		common.HandleDatabaseError(w, tx.Error)
		return
	}

	common.SendResponse(w, http.StatusOK, dbTable)
}

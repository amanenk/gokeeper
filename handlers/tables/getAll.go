package tables

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/models/table"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var tables []table.Table
	isBusy := r.URL.Query().Get("isBusy")

	chain := database.DB.Where("")
	if isBusy == "true" || isBusy == "false" {
		boolValue := isBusy == "true"
		chain.Where("is_busy = ?", boolValue)
	}

	if err := chain.Find(&tables).Error; err != nil {
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendResponse(w, http.StatusOK, tables)
}

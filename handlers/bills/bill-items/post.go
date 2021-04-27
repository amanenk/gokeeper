package bill_items

import (
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/bill"
	ordered_meal "github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func Post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderObj, err := common.GetOrderEditableByWaiter(w, r)
	if err != nil {
		logger.WithCtxValue(r.Context()).Error("problems getting users order", zap.Error(err))
	}

	orderedItemId, ok := vars["orderedItemId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("orderedItemId"))
		return
	}

	billId, ok := vars["billId"]
	if !ok {
		logger.Get().Error("missing parameter")
		common.SendError(w, errorTypes.NewNoFieldError("billId"))
		return
	}

	var billObj bill.Bill
	if err := database.Get().Where("order_id = ?", orderObj.ID).First(&billObj, billId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	var orderedMealObj ordered_meal.OrderedMeal
	if err := database.Get().First(&orderedMealObj, orderedItemId).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	uint64BillId, _ := strconv.ParseUint(billId, 10, 32)
	uintBillId := uint(uint64BillId)
	orderedMealObj.BillID = &uintBillId

	if err := database.Get().Save(orderedMealObj).Error; err != nil {
		logger.WithCtxValue(r.Context()).Error("database error", zap.Error(err))
		common.HandleDatabaseError(w, err)
		return
	}

	common.SendOk(w)
}

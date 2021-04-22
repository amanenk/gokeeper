package common

import (
	"errors"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"gorm.io/gorm"
	"net/http"
)

func HandleDatabaseError(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Get().Error("does not exist")
		//todo return error
	} else {
		logger.Get().Error("some other error")
		//todo return error
	}
	SendResponse(w, http.StatusInternalServerError, "something wrong") // todo remove it and use errors below
}

// SendExtError sends extError with code selected based on the ErrorCode.
func SendError(w http.ResponseWriter, error errorTypes.ApiError) {
	SendResponse(w, error.Code, error)
}

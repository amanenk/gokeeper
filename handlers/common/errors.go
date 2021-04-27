package common

import (
	"errors"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"gorm.io/gorm"
	"net/http"
)

func HandleDatabaseError(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		SendError(w, errorTypes.NewDBNotFoundError())
		return
	}
	SendError(w, errorTypes.NewDBGeneralError())
}

func HandleValidationError(w http.ResponseWriter, err error) {
	//todo add parsing of validation error
	SendError(w, errorTypes.NewBadRequestError("failed to validate json"))
}

// SendExtError sends extError with code selected based on the ErrorCode.
func SendError(w http.ResponseWriter, error errorTypes.ApiError) {
	SendResponse(w, error.Code, error)
}

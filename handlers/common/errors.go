package common

import (
	"errors"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"gorm.io/gorm"
	"net/http"
)

func HandleDatabaseError(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//todo return error
	} else {
		//todo return error
	}
	SendError(w, errorTypes.NewInternalServerError())
}

func HandleValidationError(w http.ResponseWriter, err error) {
	//todo add parsing of validation error
	SendError(w, errorTypes.NewBadRequestError("failed to validate json"))
}

// SendExtError sends extError with code selected based on the ErrorCode.
func SendError(w http.ResponseWriter, error errorTypes.ApiError) {
	SendResponse(w, error.Code, error)
}

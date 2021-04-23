package errorTypes

import "net/http"

type ApiError struct {
	Code    int    `json:"-"`
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

func NewNoFieldError(fieldName string) ApiError {
	return ApiError{http.StatusBadRequest, "field is not provided", fieldName}
}

func NewBadRequestError(message string) ApiError {
	return ApiError{http.StatusBadRequest, message, ""}
}

func NewForbiddenError() ApiError {
	return ApiError{http.StatusForbidden, "forbidden 403", ""}
}

func NewJSONParseError() ApiError {
	return ApiError{http.StatusBadRequest, "failed to parse JSON", ""}
}

func NewInternalServerError() ApiError {
	return ApiError{http.StatusInternalServerError, "internal server error", ""}
}

func NewUnauthorized() ApiError {
	return ApiError{http.StatusUnauthorized, "unauthorized", ""}
}

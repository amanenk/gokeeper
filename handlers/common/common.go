package common

import (
	"encoding/json"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

// SendResponse - encode response to json and send it.
func SendResponse(w http.ResponseWriter, statusCode int, respBody interface{}) {
	binRespBody, err := json.Marshal(respBody)
	if err != nil {
		logger.Get().Error("failed to marshal response body to json", zap.Error(err))
		statusCode = http.StatusInternalServerError
	}

	SendRawResponse(w, statusCode, binRespBody)
}

func SendOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

// SendRawResponse sends any raw ([]byte) response.
func SendRawResponse(w http.ResponseWriter, statusCode int, binBody []byte) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	w.WriteHeader(statusCode)
	_, err := w.Write(binBody)
	if err != nil {
		logger.Get().Error("failed to write response body", zap.Error(err))
	}
}

// UnmarshalRequestBody unmarshals request body.
func UnmarshalRequestBody(r *http.Request, body interface{}) *errorTypes.ApiError {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serverError := errorTypes.NewBadRequestError("failed to read JSON body")
		logger.WithCtxValue(r.Context()).Error("Invalid request body", zap.Error(err))
		return &serverError
	}
	defer r.Body.Close()

	err = json.Unmarshal(reqBody, body)
	if err != nil {
		serverError := errorTypes.NewBadRequestError("parse JSON body")
		logger.WithCtxValue(r.Context()).Error("Invalid JSON request body",
			zap.Error(err), zap.String("Corrupted JSON", string(reqBody)))

		return &serverError
	}

	return nil
}

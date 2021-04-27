package middlewares

import (
	"github.com/fdistorted/gokeeper/context"
	"github.com/google/uuid"
	"net/http"
)

const requestIDHeader = "X-Request-Id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			next.ServeHTTP(w, r.WithContext(context.WithRequestID(ctx, requestID)))
		},
	)
}

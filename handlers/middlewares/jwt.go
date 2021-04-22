package middlewares

import (
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/handlers/middlewares/role"
	"github.com/fdistorted/gokeeper/logger"
	"go.uber.org/zap"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get(authorizationHeader)
			logger.WithCtxValue(r.Context()).Debug("got token", zap.String("token", token))
			//todo decode token and handle its roles
			test := "admin" // todo take it from token

			allowedRoles := role.GetAllowedRoles(r.Context())

			//logger.WithCtxValue(r.Context()).Debug("allowedRoles", zap.Strings("roles", allowedRoles))

			for _, role := range allowedRoles {
				if test == string(role) {
					next.ServeHTTP(w, r)
					return
				}
			}
			common.SendError(w, errorTypes.NewForbiddenError())
		},
	)
}

package middlewares

import (
	"github.com/fdistorted/gokeeper/handlers/common"
	"github.com/fdistorted/gokeeper/handlers/common/errorTypes"
	"github.com/fdistorted/gokeeper/handlers/middlewares/role"
	"github.com/fdistorted/gokeeper/jwt"
	"github.com/fdistorted/gokeeper/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func JWT(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get(authorizationHeader)
			if authHeader == "" {
				common.SendError(w, errorTypes.NewUnauthorized())
				return
			}
			logger.WithCtxValue(r.Context()).Debug("got token", zap.String("token", authHeader))
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 {
				common.SendError(w, errorTypes.NewUnauthorized())
				return
			}
			claims, err := jwt.VerifyToken(tokenParts[1])
			if err != nil {
				common.SendError(w, errorTypes.NewUnauthorized())
				return
			}

			allowedRoles := role.GetAllowedRoles(r.Context())
			for _, role := range allowedRoles {
				if claims.Audience == string(role) {
					next.ServeHTTP(w, r)
					return
				}
			}
			common.SendError(w, errorTypes.NewForbiddenError())
		},
	)
}

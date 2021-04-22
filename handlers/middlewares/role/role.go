package role

import (
	"context"
	"net/http"
)

const (
	allowedRolesKey = "allowed_roles"
)

type RoleFilter struct {
	roles []RoleName
}

type RoleName string

const (
	Waiter RoleName = "waiter"
	Admin  RoleName = "admin"
)

func NewRoleFilter(roles ...RoleName) *RoleFilter {
	return &RoleFilter{roles: roles}
}

func (rl *RoleFilter) Attach(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, allowedRolesKey, rl.roles)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func GetAllowedRoles(ctx context.Context) []RoleName {
	value, ok := ctx.Value(allowedRolesKey).([]RoleName)
	if !ok {
		return nil
	}

	return value
}

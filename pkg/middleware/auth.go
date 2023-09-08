package middleware

import (
	"context"
	"net/http"

	"github.com/ell1jah/db_cp/pkg/logger"
)

const sessionHeader = "Authorization"

type AuthSessionsManager interface {
	GetUser(string) (int, string, error)
}

type AuthContextManager interface {
	ContextWithUserID(context.Context, int) context.Context
}

type AuthHandler struct {
	SessionManager AuthSessionsManager
	Logger         logger.Logger
	ContextManager AuthContextManager
}

func (au *AuthHandler) Auth(next http.Handler, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(sessionHeader)
		if token == "" {
			au.Logger.Infow("authoriztion",
				"url", r.URL.Path,
				"method", r.Method,
				"remote_addr", r.RemoteAddr,
				"auth result", "session header not found")

			http.Error(w, "no auth", http.StatusUnauthorized)
			return
		}

		userID, userRole, err := au.SessionManager.GetUser(token)
		if err != nil {
			au.Logger.Infow("authoriztion",
				"url", r.URL.Path,
				"method", r.Method,
				"remote_addr", r.RemoteAddr,
				"auth result", "user not found",
				"GetUser error", err)

			http.Error(w, "no auth", http.StatusUnauthorized)
			return
		}

		if len(roles) > 0 {
			roleMatch := false
			for _, role := range roles {
				if userRole == role {
					roleMatch = true
					break
				}
			}
			if !roleMatch {
				au.Logger.Infow("authoriztion",
					"url", r.URL.Path,
					"method", r.Method,
					"remote_addr", r.RemoteAddr,
					"auth result", "user role doesn`t match",
					"userID", userID,
					"userRole", userRole)
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}

		au.Logger.Infow("authoriztion",
			"url", r.URL.Path,
			"method", r.Method,
			"remote_addr", r.RemoteAddr,
			"auth result", "success",
			"userID", userID,
			"userRole", userRole)

		ctx := au.ContextManager.ContextWithUserID(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

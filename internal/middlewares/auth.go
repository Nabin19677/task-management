package middlewares

import (
	"context"
	"net/http"

	"anilkhadka.com.np/task-management/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the presence of an authentication token
		authTokenCookie, err := r.Cookie("auth_token")

		if err != nil || authTokenCookie == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		claims, err := utils.ParseJWTToken(authTokenCookie.Value)
		if err != nil {
			return
		}

		// Access user_id (jti) from the token claims
		userID, ok := claims["jti"].(string)
		if !ok {
			return
		}

		// Add user_id to the request context
		ctx := context.WithValue(r.Context(), "user_id", userID)
		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

	})
}

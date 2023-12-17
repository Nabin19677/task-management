package middlewares

import (
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the presence of an authentication token
		authTokenCookie, err := r.Cookie("auth_token")

		if err != nil || authTokenCookie == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// TODO: Validate the authentication token
		// You need to implement token validation (e.g., verify signature, check expiration) here.

		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

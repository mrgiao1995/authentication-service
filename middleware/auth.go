package middleware

import (
	"context"
	"net/http"
	"strings"
)

func HttpAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		h := r.Header.Get("Authorization")
		// Missing header or invalid header
		if h == "" || len(strings.Split(h, " ")) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed token"))
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		jwtToken := string(h[1])
		ctx = context.WithValue(r.Context(), "accessToken", jwtToken)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

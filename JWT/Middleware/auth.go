package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super_secret_key")

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Missing token", 401)
			return
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", 401)
			return
		}

		next(w, r)
	}
}

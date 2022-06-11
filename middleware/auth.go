package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func WithToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		_, err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		ctx := context.WithValue(r.Context(), "token", tokenString)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// JWT Token 検証
func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

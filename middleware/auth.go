package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Key int

const (
	UserKey Key = iota
)

func WithToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		userID, err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		// ユーザーIDをコンテキストに渡す
		ctx := context.WithValue(r.Context(), UserKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// JWT Token 検証
func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return "", err
	}
	// JWTのペイロードからユーザーIDを取得
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["iss"].(string)
	return userID, nil
}

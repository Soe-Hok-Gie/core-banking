package middleware

import "net/http"

type ContextKey string

const UserIDKey ContextKey = "userID"

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		})
	}
}

package middlewares

import (
	"errors"
	"net/http"

	"github.com/luizclaudioholanda/go-rest-app/api/auth"
	"github.com/luizclaudioholanda/go-rest-app/api/responses"
)

// SetMiddlewareJSON Set json middleware
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication Set auth middleware
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

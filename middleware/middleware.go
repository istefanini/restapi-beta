package middleware

import (
	"net/http"
	"github.com/google/uuid"
)

func requestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, r)
	})	
}
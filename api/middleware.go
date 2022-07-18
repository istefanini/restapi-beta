package api

import (
	"net/http"
)

var api_key = "sOvGUDFI7BaEhhNKdijp3cyta6Kbvxc5"

func requestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("x-api-key")
		if api_key!=requestID{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header().Set("x-api-key", requestID)
		next.ServeHTTP(w, r)
	})	
}

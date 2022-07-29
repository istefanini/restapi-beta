package api

import (
	"encoding/json"
	"net/http"
)

var api_key = "sOvGUDFI7BaEhhNKdijp3cyta6Kbvxc5"

func requestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("x-api-key")
		if api_key != requestID {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			resp := make(map[string]string)
			resp["Message"] = "Unauthorized access without api key"
			jsonResp, _ := json.Marshal(resp)
			w.Write(jsonResp)
			return
		}
		w.Header().Set("x-api-key", requestID)
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

const apiKeyHeader = "X-API-Key"
const validKey = "secret123"

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log method & path
		log.Printf("%s %s", r.Method, r.URL.Path)

		key := r.Header.Get(apiKeyHeader)
		if key != validKey {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

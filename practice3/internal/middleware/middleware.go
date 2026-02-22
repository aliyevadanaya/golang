package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf(
			"Time: %s | Method: %s | Path: %s",
			start.Format(time.RFC3339),
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)
	})
}

func APIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var validAPIKey = os.Getenv("API_KEY")
		key := r.Header.Get("X-API-KEY")

		if key != validAPIKey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"error":"unauthorized"}`))
			return
		}

		next.ServeHTTP(w, r)
	})
}

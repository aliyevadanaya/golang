package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Printf(
			"%s %s %s\n",
			time.Now().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)
	})
}

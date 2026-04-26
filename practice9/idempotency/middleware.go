package idempotency

import (
	"net/http"
	"net/http/httptest"
)

func Middleware(store *Store, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		key := r.Header.Get("Idempotency-Key")
		if key == "" {
			http.Error(w, "Idempotency-Key required", http.StatusBadRequest)
			return
		}

		if cached, exists := store.Get(key); exists {
			if cached.Completed {
				w.WriteHeader(cached.StatusCode)
				w.Write(cached.Body)
			} else {
				http.Error(w, "Duplicate in progress", http.StatusConflict)
			}
			return
		}

		if !store.Start(key) {
			http.Error(w, "Duplicate in progress", http.StatusConflict)
			return
		}

		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		store.Finish(key, rec.Code, rec.Body.Bytes())

		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())
	})
}

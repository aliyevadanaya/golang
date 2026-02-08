package main

import (
	"fmt"
	"net/http"

	"practice2/internal/handlers"
	"practice2/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			handlers.GetTasks(w, r)
			return
		}

		if r.Method == http.MethodPost {
			handlers.CreateTask(w, r)
			return
		}

		if r.Method == http.MethodPatch {
			handlers.UpdateTask(w, r)
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	})
	
	handler := middleware.Logger(
		middleware.Auth(mux),
	)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", handler)
}

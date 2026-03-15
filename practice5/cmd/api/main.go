package main

import (
	"net/http"
	"practice5/internal/database"
	"practice5/internal/handler"
	"practice5/internal/repository"
)

func main() {
	db := database.InitPostgres()
	mux := http.NewServeMux()
	repo := repository.NewRepository(db)
	handler := handler.NewUserHandler(repo)

	mux.HandleFunc("/users", handler.GetUsers)
	mux.HandleFunc("/users/common", handler.GetCommonFriends)

	http.ListenAndServe(":8080", mux)
}

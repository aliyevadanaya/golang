package handler

import (
	"encoding/json"
	"net/http"
	"practice5/internal/repository"
	"strconv"
)

type UserHandler struct {
	repo *repository.Repository
}

func NewUserHandler(repo *repository.Repository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	orderBy := r.URL.Query().Get("order_by")
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")
	gender := r.URL.Query().Get("gender")

	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	if page == 0 {
		page = 1
	}

	if pageSize == 0 {
		pageSize = 5
	}

	result, err := h.repo.GetPaginatedUsers(page, pageSize, orderBy, name, email, gender)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *UserHandler) GetCommonFriends(w http.ResponseWriter, r *http.Request) {

	user1Str := r.URL.Query().Get("user1")
	user2Str := r.URL.Query().Get("user2")

	user1, _ := strconv.Atoi(user1Str)
	user2, _ := strconv.Atoi(user2Str)

	friends, err := h.repo.GetCommonFriends(user1, user2)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

package handler

import (
	"encoding/json"
	"net/http"
	"practice3/internal/usecase"
	"practice3/pkg/modules"
	"strconv"
)

type Handler struct {
	usecase *usecase.UserUsecase
}

func NewHandler(u *usecase.UserUsecase) *Handler {
	return &Handler{u}
}

// Health godoc
// @Summary Health check
// @Tags System
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// GetUsers godoc
// @Summary Get all users
// @Description Get list of active users with pagination
// @Tags Users
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} modules.User
// @Security ApiKeyAuth
// @Router /users [get]
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10 // по дефолту
	offset := 0

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil {
			limit = l
		}
	}

	if offsetStr != "" {
		o, err := strconv.Atoi(offsetStr)
		if err == nil {
			offset = o
		}
	}

	users, err := h.usecase.GetUsers(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get single active user by ID
// @Tags Users
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} modules.User
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /user [get]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user modules.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	id, err := h.usecase.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get single active user by ID
// @Tags Users
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} modules.User
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /user [get]
func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)
	//fmt.Println("ID from request:", id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// UpdateUser godoc
// @Summary Update user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body modules.User true "Updated User"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users/update [put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user modules.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	err = h.usecase.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteUser godoc
// @Summary Soft delete user
// @Tags Users
// @Produce json
// @Param id query int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /users/delete [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
	}

	_, err = h.usecase.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	//json.NewEncoder(w).Encode(map[string]string{"message": "great Danaya, user deleted!"})
	w.WriteHeader(http.StatusOK)
}

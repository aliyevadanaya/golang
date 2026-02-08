package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"practice2/internal/models"
)

var tasks = make(map[int]models.Task)
var nextID = 1

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")

	if idParam != "" {
		id, err := strconv.Atoi(idParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
			return
		}
		task, ok := tasks[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
			return
		}

		json.NewEncoder(w).Encode(task)
		return

	}
	allTasks := make([]models.Task, 0)

	for _, task := range tasks {
		allTasks = append(allTasks, task)
	}

	json.NewEncoder(w).Encode(allTasks)

}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input struct {
		Title string `json:"title"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid title"})
		return
	}

	task := models.Task{
		ID:    nextID,
		Title: input.Title,
		Done:  false,
	}

	tasks[nextID] = task
	nextID++

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := r.URL.Query().Get("id")
	if idParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid id",
		})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid id",
		})
		return
	}

	var input struct {
		Done bool `json:"done"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid body",
		})
		return
	}

	task, ok := tasks[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "task not found",
		})
		return
	}

	task.Done = input.Done
	tasks[id] = task

	json.NewEncoder(w).Encode(map[string]bool{
		"updated": true,
	})
}

package handlers

import (
	"ecom_project/internal/domain/todos"
	"encoding/json"
	"net/http"
	"slices"
	"strconv"

	"github.com/go-chi/chi/v5"
)

var taskList = []todos.Todos{
	{
		Id:          1,
		Title:       "First Task",
		Description: "Join in a meeting",
		IsCompleted: false,
	},
	{
		Id:          2,
		Title:       "Second Task",
		Description: "Join in a meeting",
		IsCompleted: false,
	},
	{
		Id:          3,
		Title:       "Third Task",
		Description: "Join in a meeting",
		IsCompleted: false,
	},
	{
		Id:          4,
		Title:       "Fourth Task",
		Description: "Join in a meeting",
		IsCompleted: false,
	},
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var newTask todos.Todos

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newTask)

	if err != nil {
		http.Error(w, "Error fetching todo: "+err.Error(), http.StatusBadRequest)
		return
	}

	newTask.Id = len(taskList) + 1

	taskList = append(taskList, newTask)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskList)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	taskId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(taskId)

	if err != nil {
		http.Error(w, "Invalid Task Id", http.StatusBadRequest)
		return
	}
	var updatedTask todos.Todos
	var found bool

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	found = false

	for i := 0; i < len(taskList); i++ {
		if taskList[i].Id == id {
			taskList[i].Title = updatedTask.Title
			taskList[i].Description = updatedTask.Description
			taskList[i].IsCompleted = updatedTask.IsCompleted

			found = true

			updatedTask = taskList[i]
			break
		}
	}

	if found == false {
		http.Error(w, "Requested Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedTask)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(taskId)

	if err != nil {
		http.Error(w, "Invalid Task Id", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(taskList); i++ {
		if taskList[i].Id == id {
			taskList = slices.Delete(taskList, i, i+1)
		}
	}

	json.NewEncoder(w).Encode(&taskList)
}

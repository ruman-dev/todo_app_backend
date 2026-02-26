package todosRoutes

import (
	"ecom_project/cmd/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"slices"
	"strconv"
)

var todos = []domain.Todos{
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

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", listTasks)
	r.Post("/", createTask)
	r.Put("/{id}", updateTask)
	r.Delete("/{id}", deleteTask)

	return r
}

func listTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func createTask(w http.ResponseWriter, r *http.Request) {

	var newTask domain.Todos

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newTask)

	if err != nil {
		http.Error(w, "Error fetching todo: "+err.Error(), http.StatusBadRequest)
		return
	}

	newTask.Id = len(todos) + 1

	todos = append(todos, newTask)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func updateTask(w http.ResponseWriter, r *http.Request) {

	taskId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(taskId)

	if err != nil {
		http.Error(w, "Invalid Task Id", http.StatusBadRequest)
		return
	}
	var updatedTask domain.Todos
	var found bool

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	found = false

	for i := 0; i < len(todos); i++ {
		if todos[i].Id == id {
			todos[i].Title = updatedTask.Title
			todos[i].Description = updatedTask.Description
			todos[i].IsCompleted = updatedTask.IsCompleted

			found = true

			updatedTask = todos[i]
			break
		}
	}

	if found == false {
		http.Error(w, "Requested Task not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedTask)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(taskId)

	if err != nil {
		http.Error(w, "Invalid Task Id", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(todos); i++ {
		if todos[i].Id == id {
			todos = slices.Delete(todos, i, i+1)
		}
	}

	json.NewEncoder(w).Encode(&todos)
}

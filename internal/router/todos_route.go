package router

import (
	"ecom_project/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func TodosRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handlers.ListTasks)
	r.Post("/", handlers.CreateTask)
	r.Put("/{id}", handlers.UpdateTask)
	r.Delete("/{id}", handlers.DeleteTask)

	return r
}

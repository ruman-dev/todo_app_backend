package main

import (
	todosRoutes "ecom_project/cmd/internal/router"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", greetingRoute)
	r.Mount("/todos", todosRoutes.Routes())

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", r)
}

func greetingRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to our server"))
}

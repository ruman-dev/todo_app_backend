package server

import (
	"ecom_project/internal/handlers"
	"ecom_project/internal/router"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func Serve() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tokenStr := router.Init()

	r.Get("/", handlers.Greeting)
	r.Mount("/auth", router.UserRoutes())

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenStr))

		r.Use(jwtauth.Authenticator(tokenStr))

		r.Mount("/todos", router.TodosRoutes())
	})

	log.Println("Server running on :3000")
	http.ListenAndServe(":3000", r)
}

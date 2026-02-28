package router

import (
	"ecom_project/internal/handlers"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

var token *jwtauth.JWTAuth

func Init() *jwtauth.JWTAuth {
	authToken := jwtauth.New("HS256", []byte("my-secret"), nil)
	_, tokenString, _ := authToken.Encode(map[string]interface{}{"user_id": 1234})

	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

	return authToken
}

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/register", handlers.RegisterUser)
	r.Post("/login", handlers.LoginUser)

	return r
}

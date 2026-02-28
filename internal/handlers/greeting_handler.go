package handlers

import "net/http"

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to our server"))
}

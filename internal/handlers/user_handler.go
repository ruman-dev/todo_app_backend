package handlers

import (
	"ecom_project/internal/domain/users"
	"encoding/json"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var newUser users.User
	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()
	err := decoder.Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	user := users.Create(newUser)

	if user == nil {
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser users.Login

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()
	decoder.Decode(&loginUser)

	loggedInUser := users.LoginUser(loginUser)

	if loggedInUser == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&loggedInUser)

}

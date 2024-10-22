package handlers

import (
	"api/auth"
	"api/model"
	"api/services"
	"encoding/json"
	"net/http"
)

type RegisterUserHandler struct {
	userService services.UserService
}

func CreateRegisterUserHandler(userService services.UserService) *RegisterUserHandler {
	return &RegisterUserHandler{userService: userService}
}

func (handler *RegisterUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	println(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	err = handler.userService.CreateUser(user)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

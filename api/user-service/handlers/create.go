package handlers

import (
	"api/model"
	"api/services"
	"encoding/json"
	"net/http"
)

type CreateUserHandler struct {
	userService services.UserService
}

func CreateCreateUserHandler(userService services.UserService) *CreateUserHandler {
	return &CreateUserHandler{userService: userService}
}

func (handler *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.userService.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User added successfully"})
}

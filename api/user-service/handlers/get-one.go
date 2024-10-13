package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/services"
	"github.com/gorilla/mux"
)

// Handler для получения пользователя по ID
type GetUserHandler struct {
	userService services.UserService
}

// Создаем новый GetUserHandler
func CreateGetUserHandler(userService services.UserService) *GetUserHandler {
	return &GetUserHandler{userService: userService}
}

// Обработчик запроса
func (handler *GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	user, err := handler.userService.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/services"
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
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	user, err := handler.userService.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"api/services"
)

// Handler для получения всех пользователей
type GetUsersHandler struct {
	userService services.UserService
}

// Создаем новый GetUsersHandler
func CreateGetUsersHandler(userService services.UserService) *GetUsersHandler {
	return &GetUsersHandler{userService: userService}
}

// Обработчик запроса
func (h *GetUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

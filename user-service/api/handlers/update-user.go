// handlers/update_user.go
package handlers

import (
	"api/models"
	"api/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UpdateUserHandler struct {
	userService services.UserService
}

func CreateUpdateUserHandler(userService services.UserService) *UpdateUserHandler {
	return &UpdateUserHandler{userService: userService}
}

// UpdateUserHandler обновляет данные пользователя по ID
func (handler *UpdateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из URL
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Чтение данных пользователя из тела запроса
	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Устанавливаем ID пользователя в структуре
	user.ID = userID

	// Обновляем пользователя через сервис
	err = handler.userService.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

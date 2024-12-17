package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"user-service/internal/auth"
	"user-service/internal/model"
	"user-service/internal/service"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	service service.UserService
}

type Credentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Claims struct {
	Name   string   `json:"name"`
	UserId int      `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

type AssignRoleRequest struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}
type RemoveRolesRequest struct {
	Roles []string `json:"roles"`
}

func CreateUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.service.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User added successfully"})
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = handler.service.DeleteUser(userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Обработчик запроса
func (handler *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Неверный формат ID", http.StatusBadRequest)
		return
	}

	user, err := handler.service.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, token, err := handler.service.Login(creds.Name, creds.Password)
	fmt.Println(err, "LOGIN")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	tokenStr := tokenString[len("Bearer "):]
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Извлекаем данные пользователя из базы по имени пользователя из токена
	user, err := handler.service.GetUserByUsername(claims.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Возвращаем профиль пользователя
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Name     string   `json:"name"`
		Email    string   `json:"email"`
		Password string   `json:"password"`
		Roles    []string `json:"roles"` // Добавляем роли
	}

	// var user model.User

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// user.Password = hashedPassword
	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	if len(request.Roles) == 0 {
		request.Roles = append(request.Roles, "buyer")
	}

	err = handler.service.Register(user, request.Roles)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из URL
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Чтение данных пользователя из тела запроса
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Устанавливаем ID пользователя в структуре
	user.ID = userID

	// Обновляем пользователя через сервис
	err = handler.service.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// change roles user

func (handler *UserHandler) AddRoleToUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tokenString := r.Header.Get("Authorization")
	tokenStr := tokenString[len("Bearer "):]
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !auth.ContainsRole(claims.Roles, "admin") {
		http.Error(w, "Forbidden: Insufficient permissions", http.StatusForbidden)
		return

	}

	// Извлекаем ID пользователя и роль из запроса
	userID := vars["id"]

	// Преобразуем userID в int
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Извлекаем данные ролей из тела запроса
	var request struct {
		Roles []string `json:"roles"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Получаем пользователя из базы данных
	user, err := handler.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}

	// Добавляем роли пользователю
	for _, roleName := range request.Roles {
		// Проверка, есть ли уже такая роль
		roleExists := false
		for _, role := range user.Roles {
			if role.Name == roleName {
				roleExists = true
				break
			}
		}

		if roleExists {
			continue // Если роль уже есть, пропускаем её
		}

		// Добавляем роль
		err = handler.service.AddRoleToUser(user, roleName)
		if err != nil {
			http.Error(w, "Failed to add role", http.StatusInternalServerError)
			return
		}
	}

	// Успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Roles added successfully"))
}

func (handler *UserHandler) RemoveRolesFromUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	tokenString := r.Header.Get("Authorization")
	tokenStr := tokenString[len("Bearer "):]
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !auth.ContainsRole(claims.Roles, "admin") {
		http.Error(w, "Forbidden: Insufficient permissions", http.StatusForbidden)
		return
	}

	// Извлекаем ID пользователя из URL
	userID := vars["id"]

	// Преобразуем userID в int
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Извлекаем тело запроса
	var request RemoveRolesRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Извлекаем данные пользователя из базы
	user, err := handler.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		return
	}

	// Удаляем роли у пользователя
	for _, roleName := range request.Roles {
		err := handler.service.RemoveRoleFromUser(user, roleName)
		if err != nil {
			http.Error(w, "Failed to remove role: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Успешный ответ
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Roles removed successfully"))
}

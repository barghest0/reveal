package handlers

import (
	"api/auth"
	"api/services"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type ProfileUserHandler struct {
	userService services.UserService
}

func CreateProfileUserHandler(userService services.UserService) *ProfileUserHandler {
	return &ProfileUserHandler{userService: userService}
}

func (handler *ProfileUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	user, err := handler.userService.GetUserByUsername(claims.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Возвращаем профиль пользователя
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

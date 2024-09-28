package handlers

import (
	"api/auth"
	"api/services"
	"encoding/json"
	"log"
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
	cookie, err := r.Cookie("token")
	log.Println(r.Cookies(), "profile handler", err)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

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

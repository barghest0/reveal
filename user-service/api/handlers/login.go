package handlers

import (
	"api/auth"
	"api/services"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type Credentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

type LoginUserHandler struct {
	userService services.UserService
}

func CreateLoginUserHandler(userService services.UserService) *LoginUserHandler {
	return &LoginUserHandler{userService: userService}
}

func (handler *LoginUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := handler.userService.Login(creds.Name, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &Claims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+tokenString)

	w.WriteHeader(http.StatusOK)
}

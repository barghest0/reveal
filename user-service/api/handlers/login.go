package handlers

import (
	"api/services"
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

var jwtKey = []byte("my_secret_key")

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

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
}

package auth

import "github.com/golang-jwt/jwt"

var JwtKey = []byte("key")

type Claims struct {
	Name   string   `json:"name"`
	UserId int      `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

func GetAuthTokenClaims(token string) (*jwt.Token, *Claims, error) {
	tokenStr := token[len("Bearer "):]
	claims := &Claims{}
	newToken, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	return newToken, claims, err
}

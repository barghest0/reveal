package routes

import (
	"user-service/handler"

	"github.com/gorilla/mux"
)

// Инициализация маршрутов
func InitRoutes(h *handler.UserHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/users", h.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")

	router.HandleFunc("/users/register", h.Register).Methods("POST")
	router.HandleFunc("/users/login", h.Login).Methods("POST")
	router.HandleFunc("/users/profile", h.GetProfile).Methods("GET")

	return router
}

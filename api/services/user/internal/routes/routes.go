package routes

import (
	"user-service/internal/handler"

	"github.com/gorilla/mux"
)

// Инициализация маршрутов
func InitRoutes(h *handler.UserHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users/register", h.Register).Methods("POST")
	router.HandleFunc("/users/login", h.Login).Methods("POST")
	router.HandleFunc("/users/profile", h.GetProfile).Methods("GET")

	router.HandleFunc("/users", h.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", h.GetUserByID).Methods("GET")
	router.HandleFunc("/users", h.CreateUser).Methods("POST")

	router.HandleFunc("/users/{id}/roles", h.AddRoleToUser).Methods("POST")
	router.HandleFunc("/users/{id}/roles", h.RemoveRolesFromUser).Methods("DELETE")

	router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")

	return router
}

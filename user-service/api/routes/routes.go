package routes

import (
	"api/handlers"
	"api/services"
	"github.com/gorilla/mux"
)

// Инициализация маршрутов
func InitRoutes(userService services.UserService) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/users", handlers.CreateGetUsersHandler(userService)).Methods("GET")
	router.Handle("/users/{id}", handlers.CreateGetUserHandler(userService)).Methods("GET")
	router.Handle("/users", handlers.CreateCreateUserHandler(userService)).Methods("POST")
	router.Handle("/users/{id}", handlers.CreateUpdateUserHandler(userService)).Methods("PUT")
	router.Handle("/users/{id}", handlers.CreateDeleteUserHandler(userService)).Methods("DELETE")

	return router
}

package routes

import (
	"api/handlers"
	"api/services"
	"github.com/gorilla/mux"
)

// Инициализация маршрутов
func InitRoutes(userService services.UserService) *mux.Router {
	router := mux.NewRouter()

	// Создаем обработчики с сервисом
	getUsersHandler := handlers.CreateGetUsersHandler(userService)
	getUserHandler := handlers.CreateGetUserHandler(userService)
	// createUserHandler := handlers.NewCreateUserHandler(userService)
	// updateUserHandler := handlers.NewUpdateUserHandler(userService)
	// deleteUserHandler := handlers.NewDeleteUserHandler(userService)

	// Назначаем маршруты
	router.Handle("/users", getUsersHandler).Methods("GET")
	router.Handle("/user", getUserHandler).Methods("GET")
	// router.Handle("/user/create", createUserHandler).Methods("POST")
	// router.Handle("/user/update", updateUserHandler).Methods("PUT")
	// router.Handle("/user/delete", deleteUserHandler).Methods("DELETE")

	return router
}

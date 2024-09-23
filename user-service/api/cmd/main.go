package main

import (
	"api/internal/config"
	"api/internal/db"
	"api/routes"
	"api/services"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Разрешить доступ с любых доменов
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Разрешенные методы
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Разрешенные заголовки

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	app_config := config.LoadConfig()
	db_config := config.LoadDBConfig()

	database, error := db.ConnectDB(db_config)

	if error != nil {
		log.Fatalf("failed to connect to the databalse: %v", error)
	}

	userService := services.CreateUserService(database)
	router := routes.InitRoutes(userService)

	log.Printf("Server starting on port %s", app_config.ServerHost+":"+app_config.Port)
	http.ListenAndServe(":"+app_config.Port, corsMiddleware(router))

}

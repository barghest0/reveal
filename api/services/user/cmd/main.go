package main

import (
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/db"
	"user-service/internal/handler"
	"user-service/internal/messaging"
	"user-service/internal/repository"
	"user-service/internal/routes"
	"user-service/internal/service"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                              // Разрешить доступ с любых доменов
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")            // Разрешенные методы
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, withCredentials") // Разрешенные заголовки
		w.Header().Set("Access-Control-Allow-Credentials", "true")

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
	rabbitmqURL := app_config.RabbitMQURL

	database, error := db.ConnectDB(db_config)

	if error != nil {
		log.Fatalf("failed to connect to the databalse: %v", error)
	}

	db.SeedRoles(database)

	rmq, err := messaging.CreatePublisherManager(rabbitmqURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	// Объявление обменника
	if err := rmq.DeclareExchange("user_events", "fanout"); err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}

	repo := repository.CreateUserRepository(database)
	src := service.CreateUserService(repo, rmq)
	h := handler.CreateUserHandler(src)

	router := routes.InitRoutes(h)

	log.Printf("Server starting on port %s", app_config.ServerHost+":"+app_config.Port)
	http.ListenAndServe(":"+app_config.Port, corsMiddleware(router))

}

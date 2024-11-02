package main

import (
	"log"
	"net/http"
	"user-service/handler"
	"user-service/internal/config"
	"user-service/internal/db"
	"user-service/rabbitmq"
	"user-service/repository"
	"user-service/routes"
	"user-service/service"
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

	rmq, err := rabbitmq.CreateRabbitMQ(rabbitmqURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}

	repo := repository.CreateUserRepository(database)
	src := service.CreateUserService(repo)
	h := handler.CreateUserHandler(src, rmq)

	router := routes.InitRoutes(h)

	log.Printf("Server starting on port %s", app_config.ServerHost+":"+app_config.Port)
	http.ListenAndServe(":"+app_config.Port, corsMiddleware(router))

}

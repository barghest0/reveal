package main

import (
	"cart-service/internal/config"
	"cart-service/internal/db"
	"cart-service/internal/handler"
	messaging "cart-service/internal/rabbitmq"
	"cart-service/internal/repository"
	"cart-service/internal/router"
	"cart-service/internal/service"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                             // Разрешить доступ с любых доменов
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")              // Разрешенные методы
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, withCredentials, Authorization") // Разрешенные заголовки
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

	database, error := db.ConnectDB(db_config)

	rmq, err := messaging.CreateRabbitMQ("amqp://guest:guest@rabbitmq:5672/")

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	if error != nil {
		log.Fatalf("falied to connect to the database: %v", error)
	}

	repo := repository.CreateCartRepository(database)
	src := service.CreateCartService(repo, rmq)
	h := handler.CreateCartHandler(src)

	go func() {
		src.CreateCartAfterRegistration()
	}()

	router := router.CreateRouter(h)

	log.Printf("Server starting on port %s", app_config.ServerHost+":"+app_config.Port)
	http.ListenAndServe(":"+app_config.Port, corsMiddleware(router))

}

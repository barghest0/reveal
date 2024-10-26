package main

import (
	"log"
	"net/http"
	"product-service/internal/config"
	"product-service/internal/db"
	"product-service/internal/handler"
	"product-service/internal/repository"
	"product-service/internal/router"
	"product-service/internal/service"

	"github.com/barghest0/reveal/api/packages/cache"
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

	database, error := db.ConnectDB(db_config)

	if error != nil {
		log.Fatalf("failed to connect to the databalse: %v", error)
	}

	redis := cache.CreateRedisClient()
	cache_src := cache.CreateCacheService(redis)

	repo := repository.CreateProductRepository(database, cache_src)
	src := service.CreateProductService(repo)
	h := handler.CreateProductHandler(src)

	router := router.CreateRouter(h)

	log.Printf("Server starting on port %s", app_config.ServerHost+":"+app_config.Port)
	http.ListenAndServe(":"+app_config.Port, corsMiddleware(router))
}

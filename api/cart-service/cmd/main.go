package main

import (
	"net/http"

	_ "github.com/lib/pq"
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
	// app_config := config.LoadConfig()
	// db_config := config.LoadDBConfig()

	// database, error := db.ConnectDB(db_config)

	// if error != nil {
	// 	log.Fatalf("Failed to conect to the databalse: %v", error)
	// }

}

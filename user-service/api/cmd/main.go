package main

import (
	"api/internal/config"
	"api/internal/db"
	"api/routes"
	"api/services"
	"log"
	"net/http"
)

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
	http.ListenAndServe(":"+app_config.Port, router)

}

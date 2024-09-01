package main

import (
	"log"
	"net/http"
	"user-management-service/internal/config"
	"user-management-service/internal/db"
	"user-management-service/internal/user"
)

func main() {
	config := config.LoadConfig()

	database, error := db.ConnectDB(&config.DBConfig)

	if error != nil {
		log.Fatalf("failed to connect to the databalse: %v", error)
	}

	defer database.Close()

	userService := user.NewService(database)

	http.HandleFunc("/users", userService.GetUsersHandler)

	log.Printf("Server starting on port %s", config.DBConfig.Port)
	log.Fatal(http.ListenAndServe(":"+config.DBConfig.Port, nil))

}

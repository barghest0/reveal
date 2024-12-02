package main

import (
	"admin-service/handler"
	"admin-service/internal/config"
	"admin-service/router"
	"admin-service/service"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	adminService := service.CreateAdminService(cfg)

	adminHandler := &handler.AdminHandler{
		Service: adminService,
	}

	r := router.CreateRouter(adminHandler)

	log.Println("Admin service running on port", cfg.AdminPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AdminPort, r))
}

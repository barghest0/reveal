package main

import (
	"cart-service/internal/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "barghest"
	dbname   = "cart_service"
)

func main() {
	// connection string
	dsn := fmt.Sprintf("host=postgres port=%d user=%s password=%s dbname=%s sslmode=disable", port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&model.Cart{}, &model.CartItem{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Successfully connected to the database")
}

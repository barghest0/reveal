package db

import (
	"cart-service/internal/config"
	"cart-service/internal/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(db_config config.DBConfig) (*gorm.DB, error) {
	// connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db_config.Host, db_config.User, db_config.Password, db_config.Name, db_config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&model.Cart{}, &model.CartProduct{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return db, nil
}

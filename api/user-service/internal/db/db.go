package db

import (
	"api/internal/config"
	"api/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(db_config config.DBConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db_config.Host, db_config.User, db_config.Password, db_config.Name, db_config.Port)

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Миграции базы данных
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return db, nil
}

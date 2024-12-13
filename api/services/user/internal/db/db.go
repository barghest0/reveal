package db

import (
	"fmt"
	"log"
	"user-service/internal/config"
	"user-service/internal/model"

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
	if err := db.AutoMigrate(&model.User{}, &model.Role{}, &model.UserRoles{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Successfully connected to the database")

	return db, nil
}

func SeedRoles(db *gorm.DB) {
	roles := []model.Role{
		{Name: "admin"},
		{Name: "seller"},
		{Name: "buyer"},
	}
	for _, role := range roles {
		// Используем FirstOrCreate для поиска и создания роли
		if err := db.Where("name = ?", role.Name).FirstOrCreate(&role).Error; err != nil {
			log.Printf("Error adding role: %v", err)
		}
	}
	// for _, role := range roles {
	// 	if err := db.Create(&role).Error; err != nil {
	// 		log.Printf("Error adding role: %v", err)
	// 	}
	// }
}

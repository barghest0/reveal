package db

import (
	"database/sql"
	"fmt"
	"user-management-service/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDB(config *config.DBConfig) (*sql.DB, error) {

	connection := fmt.Sprintf("user=%s password=%s dbname=%s port=5432 host=%s sslmode=disable", config.User, config.Password, config.Name, config.Host)
	// connection := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", config.User, config.Password, config.Host, config.Name)

	db, error := sql.Open("postgres", connection)

	if error != nil {
		return nil, error
	}

	if error := db.Ping(); error != nil {
		return nil, error
	}

	return db, nil

}

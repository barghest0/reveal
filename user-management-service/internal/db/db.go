package db

import (
	"database/sql"
	"fmt"
	"log"
	"user-management-service/internal/config"

	_ "github.com/lib/pq"
)

func ConnectDB(config *config.DBConfig) (*sql.DB, error) {

	connection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", config.User, config.Password, config.Name, config.Host)
	// connection := fmt.Sprintf("postgres://postgres:barghest@localhost:5432/users?sslmode=disable", config.User, config.Password, config.Host, config.Name)

	db, error := sql.Open("postgres", connection)

	if error != nil {
		log.Fatal(error)
	}

	error = db.Ping()
	if error != nil {
		log.Fatal("Unable to connect to the database: ", error)
	}
	log.Println("Successfully connected to the database")

	return db, nil

}

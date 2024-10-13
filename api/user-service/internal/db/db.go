package db

import (
	"api/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB(db_config config.DBConfig) (*sql.DB, error) {

	connection := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", db_config.User, db_config.Password, db_config.Name, db_config.Host)

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

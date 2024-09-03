package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

// Handler для получения всех пользователей
func (service *Service) GetUsersHandler(writer http.ResponseWriter, reader *http.Request) {
	if reader.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := service.db.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(users)
}

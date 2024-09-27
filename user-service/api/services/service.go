package services

import (
	"api/auth"
	"api/models"
	"database/sql"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	CreateUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(id int) error
	Login(name string, password string) (models.User, error)
}

type userService struct {
	db *sql.DB
}

func CreateUserService(db *sql.DB) UserService {
	return &userService{db: db}
}

func (service *userService) GetAllUsers() ([]models.User, error) {
	return models.GetUsers(service.db)
}

func (service *userService) GetUserByID(id int) (models.User, error) {
	return models.GetUserByID(service.db, id)
}

// Создание нового пользователя
func (service *userService) CreateUser(user models.User) error {
	return models.CreateUser(service.db, user)
}

// Обновление данных пользователя
func (service *userService) UpdateUser(user models.User) error {
	return models.UpdateUser(service.db, user)
}

// Удаление пользователя
func (service *userService) DeleteUser(id int) error {
	return models.DeleteUser(service.db, id)
}

func (service *userService) Login(name string, password string) (models.User, error) {
	user, err := models.GetUserByUsername(service.db, name)
	if err != nil {
		return models.User{}, err
	}

	if auth.CheckPasswordHash(password, user.Password) {
		return models.User{}, sql.ErrNoRows
	}

	return user, nil
}

package services

import (
	"api/auth"
	"api/model"
	"database/sql"

	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
	GetUserByUsername(names string) (model.User, error)
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
	DeleteUser(id int) error
	Login(name string, password string) (model.User, error)
}

type userService struct {
	db *gorm.DB
}

func CreateUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (service *userService) GetAllUsers() ([]model.User, error) {
	return model.GetUsers(service.db)
}

func (service *userService) GetUserByID(id int) (model.User, error) {
	return model.GetUserByID(service.db, id)
}

// Создание нового пользователя
func (service *userService) CreateUser(user model.User) error {
	return model.CreateUser(service.db, &user)
}

// Обновление данных пользователя
func (service *userService) UpdateUser(user model.User) error {
	return model.UpdateUser(service.db, user)
}

// Удаление пользователя
func (service *userService) DeleteUser(id int) error {
	return model.DeleteUser(service.db, id)
}

func (service *userService) GetUserByUsername(name string) (model.User, error) {
	return model.GetUserByUsername(service.db, name)
}

func (service *userService) Login(name string, password string) (model.User, error) {
	user, err := model.GetUserByUsername(service.db, name)
	if err != nil {
		return model.User{}, err
	}

	if auth.CheckPasswordHash(password, user.Password) {
		return model.User{}, sql.ErrNoRows
	}

	return user, nil
}

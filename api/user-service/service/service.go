package service

import (
	"database/sql"
	"user-service/auth"
	"user-service/model"
	"user-service/repository"
)

type UserService interface {
	GetAllUsers() (*[]model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByUsername(names string) (*model.User, error)
	CreateUser(user model.User) error
	UpdateUser(user model.User) error
	DeleteUser(id int) error
	Login(name string, password string) (model.User, error)
}

type userService struct {
	repository repository.UserRepository
}

func CreateUserService(repository repository.UserRepository) UserService {
	return &userService{repository}
}

func (service *userService) GetAllUsers() (*[]model.User, error) {
	return service.repository.GetAll()
}

func (service *userService) GetUserByID(id int) (*model.User, error) {
	return service.repository.GetByID(id)
}

// Создание нового пользователя
func (service *userService) CreateUser(user model.User) error {
	return service.repository.Create(&user)
}

// Обновление данных пользователя
func (service *userService) UpdateUser(user model.User) error {
	return service.repository.Update(&user)
}

// Удаление пользователя
func (service *userService) DeleteUser(id int) error {
	return service.repository.Delete(id)
}

func (service *userService) GetUserByUsername(name string) (*model.User, error) {
	return service.repository.GetByUsername(name)
}

func (service *userService) Login(name string, password string) (model.User, error) {
	user, err := service.repository.GetByUsername(name)
	if err != nil {
		return model.User{}, err
	}

	if auth.CheckPasswordHash(password, user.Password) {
		return model.User{}, sql.ErrNoRows
	}

	return *user, nil
}

package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"user-service/auth"
	"user-service/messaging"
	"user-service/model"
	"user-service/repository"
)

type UserService interface {
	GetAllUsers() (*[]model.User, error)
	GetUserByID(id int) (*model.User, error)
	GetUserByUsername(names string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user model.User) error
	DeleteUser(id int) error
	Login(name string, password string) (model.User, error)
	Register(user model.User) error
}

type userService struct {
	repository repository.UserRepository
	rmq        messaging.RabbitMQ
}

func CreateUserService(repository repository.UserRepository, rmq *messaging.RabbitMQ) UserService {
	return &userService{repository, *rmq}
}

func (service *userService) GetAllUsers() (*[]model.User, error) {
	return service.repository.GetAll()
}

func (service *userService) GetUserByID(id int) (*model.User, error) {
	return service.repository.GetByID(id)
}

// Создание нового пользователя
func (service *userService) CreateUser(user *model.User) error {
	return service.repository.Create(user)
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

	if !auth.CheckPasswordHash(password, user.Password) {
		return model.User{}, sql.ErrNoRows
	}

	return *user, nil
}

type UserData struct {
	Id uint `json:"id"`
}

func (service *userService) Register(user model.User) error {

	err := service.CreateUser(&user)
	if err != nil {
		return err
	}

	message, err := json.Marshal(UserData{Id: uint(user.ID)})

	log.Printf("PUBLISH MESSAGE: %s", []byte(fmt.Sprintf("%v", message)))
	log.Printf("PUBLISH MESSAGE JSON: %s", message)

	if err != nil {
		log.Fatalf("Failed to marshal body: %v", err)
	}

	var userData UserData
	json.Unmarshal([]byte(message), &userData)

	log.Printf("USER DATA: %v", userData)

	err = service.rmq.Publish("user.registered", message)

	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	return nil
}

package service

import (
	"database/sql"
	"encoding/json"
	"log"
	"user-service/internal/auth"
	"user-service/internal/messaging"
	"user-service/internal/model"
	"user-service/internal/repository"
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
	repository       repository.UserRepository
	publisherManager messaging.PublisherManager
}

func CreateUserService(repository repository.UserRepository, publisherManager *messaging.PublisherManager) UserService {
	return &userService{repository, *publisherManager}
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

	// Публикуем событие о создании пользователя
	event := map[string]interface{}{
		"user_id": user.ID,
	}
	eventBody, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = service.publisherManager.Publish("user_events", "", eventBody)
	if err != nil {
		return err
	}

	log.Println("User creation event published successfully.")
	return nil
}

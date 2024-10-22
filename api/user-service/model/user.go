package model

import (
	"errors"
	"gorm.io/gorm"
)

// User структура для представления пользователя
type User struct {
	ID       int   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

// Получение всех пользователей
func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Получение пользователя по ID
func GetUserByID(db *gorm.DB, id int) (User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("пользователь не найден")
		}
		return user, err
	}
	return user, nil
}

// Получение пользователя по имени
func GetUserByUsername(db *gorm.DB, name string) (User, error) {
	var user User
	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("пользователь не найден")
		}
		return user, err
	}
	return user, nil
}

// Создание нового пользователя
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// Обновление данных пользователя
func UpdateUser(db *gorm.DB, user User) error {
	return db.Save(&user).Error
}

// Удаление пользователя
func DeleteUser(db *gorm.DB, id int) error {
	return db.Delete(&User{}, id).Error
}


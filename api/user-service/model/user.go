package model

// User структура для представления пользователя
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null`
	Email    string `json:"email" gorm:"unique;not null`
	Password string `json:"password" gorm:"not null`
}

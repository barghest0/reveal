package model

// User структура для представления пользователя
type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null`
	Email    string `json:"email" gorm:"unique;not null`
	Password string `json:"password" gorm:"not null`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles;"`
}

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique;not null;column:name"`
}

type UserRoles struct {
	UserID uint `gorm:"primaryKey;column:user_id"`
	RoleID uint `gorm:"primaryKey;column:role_id"`
	User   User `gorm:"foreignkey:UserID"`
	Role   Role `gorm:"foreignkey:RoleID"`
}

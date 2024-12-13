package repository

import (
	"errors"
	"fmt"
	"user-service/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetAll() (*[]model.User, error)
	GetByID(id int) (*model.User, error)
	GetByUsername(name string) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
	GetRoleByName(roleName string) (model.Role, error)

	AssociateRoles(user *model.User, roles []model.Role) error
	GetRolesForUser(user *model.User, roles *[]model.Role) error
}

type userRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAll() (*[]model.User, error) {

	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *userRepository) GetByID(id int) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, errors.New("пользователь не найден")
		}
		return &user, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(name string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, errors.New("пользователь не найден")
		}
		return &user, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(&user).Error

}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) GetRoleByName(name string) (model.Role, error) {
	var role model.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (r *userRepository) AssociateRoles(user *model.User, roles []model.Role) error {
	if err := r.db.Model(&user).Association("Roles").Append(roles); err != nil {
		return fmt.Errorf("could not associate roles: %v", err)
	}
	return nil
}

func (r *userRepository) GetRolesForUser(user *model.User, roles *[]model.Role) error {
	// Используем ассоциацию для получения ролей пользователя
	if err := r.db.Preload("Roles").First(user).Error; err != nil {
		return fmt.Errorf("could not fetch roles: %v", err)
	}
	*roles = user.Roles
	return nil
}

package services

import (
	"errors"

	"../models"
	"gorm.io/gorm"
)

// UserService 用户服务接口
type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(id uint, user *models.User) error
	DeleteUser(id uint) error
}

// userService 用户服务实现
type userService struct {
	db *gorm.DB
}

// NewUserService 创建用户服务实例
func NewUserService() UserService {
	return &userService{
		db: models.GetDB(),
	}
}

// GetAllUsers 获取所有用户
func (s *userService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID 根据ID获取用户
func (s *userService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建用户
func (s *userService) CreateUser(user *models.User) error {
	return s.db.Create(user).Error
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(id uint, user *models.User) error {
	// 检查用户是否存在
	_, err := s.GetUserByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return nil
	}

	// 更新用户信息
	user.ID = id
	return s.db.Save(user).Error
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}
package services

import (
	"testing"

	"../models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUserService(t *testing.T) {
	// 创建测试数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	assert.NoError(t, err)

	// 替换全局DB实例
	models.DB = db

	// 创建用户服务实例
	service := NewUserService()

	// 测试创建用户
	t.Run("CreateUser", func(t *testing.T) {
		user := &models.User{
			Name:  "Test User",
			Email: "test@example.com",
			Age:   25,
			Active: true,
		}

		err := service.CreateUser(user)
		assert.NoError(t, err)
		assert.NotZero(t, user.ID)
	})

	// 测试获取所有用户
	t.Run("GetAllUsers", func(t *testing.T) {
		users, err := service.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 1)
	})

	// 测试根据ID获取用户
	t.Run("GetUserByID", func(t *testing.T) {
		user, err := service.GetUserByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Test User", user.Name)

		// 测试获取不存在的用户
		user, err = service.GetUserByID(999)
		assert.NoError(t, err)
		assert.Nil(t, user)
	})

	// 测试更新用户
	t.Run("UpdateUser", func(t *testing.T) {
		updatedUser := &models.User{
			Name:  "Updated User",
			Email: "updated@example.com",
			Age:   30,
		}

		err := service.UpdateUser(1, updatedUser)
		assert.NoError(t, err)

		// 验证更新
		user, err := service.GetUserByID(1)
		assert.NoError(t, err)
		assert.Equal(t, "Updated User", user.Name)
		assert.Equal(t, "updated@example.com", user.Email)
		assert.Equal(t, 30, user.Age)
	})

	// 测试删除用户
	t.Run("DeleteUser", func(t *testing.T) {
		err := service.DeleteUser(1)
		assert.NoError(t, err)

		// 验证删除
		users, err := service.GetAllUsers()
		assert.NoError(t, err)
		assert.Len(t, users, 0)
	})
}
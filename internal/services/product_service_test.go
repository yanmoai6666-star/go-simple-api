package services

import (
	"testing"

	"../models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductService(t *testing.T) {
	// 创建测试数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	assert.NoError(t, err)

	// 替换全局DB实例
	models.DB = db

	// 创建产品服务实例
	service := NewProductService()

	// 创建测试用户
	user := &models.User{
		Name:  "Test User",
		Email: "test@example.com",
		Age:   25,
		Active: true,
	}
	db.Create(user)

	// 测试创建产品
	t.Run("CreateProduct", func(t *testing.T) {
		product := &models.Product{
			Name:        "Test Product",
			Description: "Test Description",
			Price:       99.99,
			Stock:       100,
			UserID:      user.ID,
		}

		err := service.CreateProduct(product)
		assert.NoError(t, err)
		assert.NotZero(t, product.ID)
	})

	// 测试获取所有产品
	t.Run("GetAllProducts", func(t *testing.T) {
		products, err := service.GetAllProducts()
		assert.NoError(t, err)
		assert.Len(t, products, 1)
	})

	// 测试根据ID获取产品
	t.Run("GetProductByID", func(t *testing.T) {
		product, err := service.GetProductByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, "Test Product", product.Name)

		// 测试获取不存在的产品
		product, err = service.GetProductByID(999)
		assert.NoError(t, err)
		assert.Nil(t, product)
	})

	// 测试根据用户ID获取产品
	t.Run("GetProductsByUserID", func(t *testing.T) {
		products, err := service.GetProductsByUserID(user.ID)
		assert.NoError(t, err)
		assert.Len(t, products, 1)
	})

	// 测试更新产品
	t.Run("UpdateProduct", func(t *testing.T) {
		updatedProduct := &models.Product{
			Name:        "Updated Product",
			Description: "Updated Description",
			Price:       149.99,
			Stock:       50,
			UserID:      user.ID,
		}

		err := service.UpdateProduct(1, updatedProduct)
		assert.NoError(t, err)

		// 验证更新
		product, err := service.GetProductByID(1)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Product", product.Name)
		assert.Equal(t, "Updated Description", product.Description)
		assert.Equal(t, 149.99, product.Price)
		assert.Equal(t, 50, product.Stock)
	})

	// 测试删除产品
	t.Run("DeleteProduct", func(t *testing.T) {
		err := service.DeleteProduct(1)
		assert.NoError(t, err)

		// 验证删除
		products, err := service.GetAllProducts()
		assert.NoError(t, err)
		assert.Len(t, products, 0)
	})
}
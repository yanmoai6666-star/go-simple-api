package services

import (
	"errors"

	"../models"
	"gorm.io/gorm"
)

// ProductService 产品服务接口
type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(id uint, product *models.Product) error
	DeleteProduct(id uint) error
	GetProductsByUserID(userID uint) ([]models.Product, error)
}

// productService 产品服务实现
type productService struct {
	db *gorm.DB
}

// NewProductService 创建产品服务实例
func NewProductService() ProductService {
	return &productService{
		db: models.GetDB(),
	}
}

// GetAllProducts 获取所有产品
func (s *productService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := s.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductByID 根据ID获取产品
func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := s.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// CreateProduct 创建产品
func (s *productService) CreateProduct(product *models.Product) error {
	return s.db.Create(product).Error
}

// UpdateProduct 更新产品
func (s *productService) UpdateProduct(id uint, product *models.Product) error {
	// 检查产品是否存在
	_, err := s.GetProductByID(id)
	if err != nil {
		return err
	}
	if product == nil {
		return nil
	}

	// 更新产品信息
	product.ID = id
	return s.db.Save(product).Error
}

// DeleteProduct 删除产品
func (s *productService) DeleteProduct(id uint) error {
	return s.db.Delete(&models.Product{}, id).Error
}

// GetProductsByUserID 根据用户ID获取产品
func (s *productService) GetProductsByUserID(userID uint) ([]models.Product, error) {
	var products []models.Product
	if err := s.db.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
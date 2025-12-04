package models

import (
	"time"

	"gorm.io/gorm"
)

// Product 产品模型
type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        string         `json:"name" gorm:"size:200;not null"`
	Description string         `json:"description" gorm:"type:text"`
	Price       float64        `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock       int            `json:"stock" gorm:"not null;default:0"`
	UserID      uint           `json:"user_id" gorm:"index;not null"`
	User        User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
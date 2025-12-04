package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"regexp"
)

// GenerateRandomID 生成随机ID
func GenerateRandomID(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidateName 验证用户名格式
func ValidateName(name string) error {
	if len(name) < 2 {
		return errors.New("name must be at least 2 characters long")
	}
	if len(name) > 100 {
		return errors.New("name must be at most 100 characters long")
	}
	return nil
}

// ValidatePrice 验证价格
func ValidatePrice(price float64) error {
	if price < 0 {
		return errors.New("price must be non-negative")
	}
	return nil
}

// ValidateStock 验证库存
func ValidateStock(stock int) error {
	if stock < 0 {
		return errors.New("stock must be non-negative")
	}
	return nil
}
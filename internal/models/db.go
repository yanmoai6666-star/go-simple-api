package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(dsn string) error {
	// 配置 GORM 日志
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(dsn), config)
	if err != nil {
		return err
	}

	// 自动迁移模型
	if err := db.AutoMigrate(&User{}, &Product{}); err != nil {
		return err
	}

	DB = db
	log.Println("Database connected and migrated successfully")
	return nil
}

// GetDB 获取数据库连接实例
func GetDB() *gorm.DB {
	return DB
}
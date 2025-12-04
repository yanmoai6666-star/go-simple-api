package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"../../internal/handlers"
	"../../internal/models"
	"../../internal/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestServer(t *testing.T) *gin.Engine {
	// 创建测试数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// 自动迁移模型
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	assert.NoError(t, err)

	// 替换全局DB实例
	models.DB = db

	// 创建服务实例
	userService := services.NewUserService()
	productService := services.NewProductService()

	// 创建处理器实例
	userHandler := handlers.NewUserHandler(userService)
	productHandler := handlers.NewProductHandler(productService)

	// 设置 Gin 模式为测试模式
	gin.SetMode(gin.TestMode)

	// 创建 Gin 引擎
	router := gin.Default()

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API 路由组
	api := router.Group("/api/v1")
	{
		// 用户路由
		users := api.Group("/users")
		{
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// 产品路由
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetAllProducts)
			products.GET("/:id", productHandler.GetProductByID)
			products.POST("", productHandler.CreateProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
	}

	// 返回测试服务器
	return router
}

func TestHealthCheck(t *testing.T) {
	// 设置测试服务器
	router := setupTestServer(t)

	// 创建请求
	req, err := http.NewRequest("GET", "/health", nil)
	assert.NoError(t, err)

	// 发送请求
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 验证响应
	assert.Equal(t, http.StatusOK, w.Code)

	// 解析响应体
	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 验证响应内容
	assert.Equal(t, "ok", response["status"])
}
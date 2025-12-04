module go-simple-api

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	gorm.io/driver/sqlite v1.5.2
	gorm.io/gorm v1.25.4
)

replace (
    go-simple-api/configs => ./configs
    go-simple-api/internal/handlers => ./internal/handlers
    go-simple-api/internal/models => ./internal/models
    go-simple-api/internal/services => ./internal/services
)

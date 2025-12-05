# Go Simple API

A simple RESTful API built with Go, Gin framework, and GORM.

## Features

- **RESTful API**: Implements standard HTTP methods for resource management
- **Database Integration**: Uses SQLite with GORM for data persistence
- **Configuration Management**: Environment-based configuration using Viper
- **Testing**: Unit tests and integration tests with Testify
- **GitHub Actions**: CI/CD pipeline for automated testing and building

## Project Structure

```
go-simple-api/
├── cmd/
│   └── api/
│       └── main.go              # API entry point
├── configs/
│   └── config.go                # Configuration management
├── internal/
│   ├── handlers/                # HTTP request handlers
│   │   ├── user_handler.go
│   │   └── product_handler.go
│   ├── models/                  # Data models
│   │   ├── db.go               # Database connection
│   │   ├── user.go
│   │   └── product.go
│   └── services/               # Business logic
│       ├── user_service.go
│       ├── user_service_test.go
│       └── product_service.go
│       └── product_service_test.go
├── pkg/
│   └── utils/                  # Utility functions
│       └── validation.go
├── tests/
│   └── integration/            # Integration tests
│       └── health_test.go
├── go.mod                      # Go module dependencies
├── go.sum                      # Dependency checksums
└── README.md                   # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or later

### Installation

1. Clone the repository:

```bash
git clone https://github.com/example/go-simple-api.git
cd go-simple-api
```

2. Install dependencies:

```bash
go mod download
```

3. Run the application:

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`.

### API Endpoints

#### Health Check
- `GET /health` - Check if the API is running

#### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get a user by ID
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/:id` - Update a user
- `DELETE /api/v1/users/:id` - Delete a user

#### Products
- `GET /api/v1/products` - Get all products
- `GET /api/v1/products/:id` - Get a product by ID
- `POST /api/v1/products` - Create a new product
- `PUT /api/v1/products/:id` - Update a product
- `DELETE /api/v1/products/:id` - Delete a product

## Testing

### Run Unit Tests

```bash
go test ./...
```

### Run Integration Tests

```bash
go test ./tests/integration/...
```

## Configuration

The application can be configured using environment variables or the `.env` file:

- `ENVIRONMENT`: Application environment (development/production)
- `PORT`: Port number for the API server
- `DATABASE_DSN`: Database connection string

## License

MIT

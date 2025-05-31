# 🪑 Reservation API

This is a simple reservation system API written in Go (Golang). It was created by me as a personal project to **learn Go** through hands-on practice. The API supports reservations of seats, with features like seat pricing, discount application based on seat count, and basic CRUD operations.

## 🚀 Features

- Reserve seats with pricing logic
- Apply discount codes automatically
- Seat and reservation management
- Gin-powered REST API
- PostgreSQL support
- Swagger documentation

## 📚 Tech Stack

- **Go**
- **Gin**
- **PostgreSQL**
- **swaggo/swag**

## 🛠️ How to run locally

#### 1. Set up environment variables

```bash
cp .env.dist .env
```

#### 2. Start the database

```bash
docker compose up -d
```

### 3. Run database migrations & seed the database

```
go run ./cmd/migrate up
go run ./cmd/migrate seed
```

#### 4. Run the application with live reloading

```bash
air
```

### 5. Access Swagger API documentation

`/api/v1/swagger/index.html`

package auth

import (
	"database/sql"
	"time"

	"github.com/GanderBite/reservation-api/internal/auth/infrastructure/handlers"
	infraRepo "github.com/GanderBite/reservation-api/internal/auth/infrastructure/repositories"
	"github.com/GanderBite/reservation-api/internal/auth/infrastructure/services"
	"github.com/GanderBite/reservation-api/internal/pkg/env"
)

type AuthModule struct {
	Handlers *handlers.AuthHandlers
}

func NewAuthModule(
	db *sql.DB,
) *AuthModule {
	repo := infraRepo.NewPostgresUsersRepository(db)
	auth := services.NewAuthService(
		env.GetEnvString("JWT_SECRET"),
		time.Duration(env.GetEnvInt("JWT_TTL")),
	)

	handlers := handlers.NewAuthHandlers(auth, repo)

	return &AuthModule{
		handlers,
	}
}

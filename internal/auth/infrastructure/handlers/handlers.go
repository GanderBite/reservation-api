package handlers

import (
	"github.com/GanderBite/reservation-api/internal/auth/model/repositories"
	"github.com/GanderBite/reservation-api/internal/auth/model/services"
	usecases "github.com/GanderBite/reservation-api/internal/auth/model/use-cases"
)

type AuthHandlers struct {
	SignUpHandler *signUpHandler
}

func NewAuthHandlers(
	auth services.AuthService,
	repo repositories.UsersRepository,
) *AuthHandlers {
	signUpUC := usecases.NewSignUpUseCase(
		repo,
		auth,
	)

	signUpHandler := newSignUpHandler(signUpUC)

	return &AuthHandlers{
		SignUpHandler: signUpHandler,
	}
}

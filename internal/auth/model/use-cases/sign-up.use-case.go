package usecases

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/auth/model/dtos"
	"github.com/GanderBite/reservation-api/internal/auth/model/entities"
	"github.com/GanderBite/reservation-api/internal/auth/model/repositories"
	"github.com/GanderBite/reservation-api/internal/auth/model/services"
)

type SignUpUseCase struct {
	auth services.AuthService
	repo repositories.UsersRepository
}

func NewSignUpUseCase(repo repositories.UsersRepository, auth services.AuthService) *SignUpUseCase {
	return &SignUpUseCase{
		auth: auth,
		repo: repo,
	}
}

func (uc *SignUpUseCase) Execute(ctx context.Context, dto *dtos.SignUpDto) (string, error) {
	existingUser, err := uc.repo.GetByEmailOrUsername(ctx, dto.Email, dto.Username)
	if err != nil {
		return "", err
	}

	if existingUser != nil {
		return "", entities.ErrUserAlreadyExists
	}

	err = uc.auth.ValidatePassword(dto.Password)
	if err != nil {
		return "", err
	}

	hashed, err := uc.auth.HashPassword(dto.Password)
	if err != nil {
		return "", err
	}

	identity := entities.NewIdentityFromDto(dto.Email, hashed)
	user := entities.NewUserFromDto(dto.Username)

	createdUserId, err := uc.repo.Insert(ctx, identity, user)
	if err != nil {
		return "", err
	}

	return uc.auth.GenerateToken(createdUserId)
}

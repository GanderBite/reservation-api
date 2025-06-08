package services

import "github.com/GanderBite/reservation-api/internal/pkg/types"

type AuthService interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password string) error
	ComparePassword(hashed, plain string) error
	GenerateToken(payload types.Id) (string, error)
}

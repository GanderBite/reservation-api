package services_test

import (
	"testing"

	"github.com/GanderBite/reservation-api/internal/auth/infrastructure/services"
	"github.com/GanderBite/reservation-api/internal/auth/model/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupAuthService() *services.AuthService {
	return services.NewAuthService("test-secret-key", 1) // 1 hour TTL
}

func TestHashAndComparePassword(t *testing.T) {
	s := setupAuthService()

	password := "StrongP@ss1"

	hashed, err := s.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)

	err = s.ComparePassword(hashed, password)
	assert.NoError(t, err)

	err = s.ComparePassword(hashed, "WrongPassword123!")
	assert.Error(t, err)
}

func TestValidatePassword(t *testing.T) {
	s := setupAuthService()

	valid := "GoodP@ss1"
	weak := "password"
	missingUpper := "weakpass1!"
	missingNumber := "NoNumbers!"
	missingSpecial := "NoSpecial1"

	assert.NoError(t, s.ValidatePassword(valid))
	assert.ErrorIs(t, s.ValidatePassword(weak), entities.ErrPasswordNotSecure)
	assert.ErrorIs(t, s.ValidatePassword(missingUpper), entities.ErrPasswordNotSecure)
	assert.ErrorIs(t, s.ValidatePassword(missingNumber), entities.ErrPasswordNotSecure)
	assert.ErrorIs(t, s.ValidatePassword(missingSpecial), entities.ErrPasswordNotSecure)
}

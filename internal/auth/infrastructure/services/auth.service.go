package services

import (
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/GanderBite/reservation-api/internal/auth/model/entities"
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type AuthService struct {
	secretKey string
	tokenTTL  time.Duration
}

func NewAuthService(secretKey string, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		secretKey: secretKey,
		tokenTTL:  tokenTTL,
	}
}

func (s *AuthService) HashPassword(password string) (string, error) {
	const cost = bcrypt.DefaultCost
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}

func (s *AuthService) ValidatePassword(password string) error {
	var hasUpper, hasNumber, hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsNumber(ch):
			hasNumber = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasNumber || !hasSpecial {
		return entities.ErrPasswordNotSecure
	}

	return nil
}

func (s *AuthService) GenerateToken(payload types.Id) (string, error) {
	claims := jwt.MapClaims{
		"sub": payload,
		"exp": time.Now().Add(s.tokenTTL * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
}

func (s *AuthService) ComparePassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

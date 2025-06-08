package entities

import "github.com/GanderBite/reservation-api/internal/pkg/types"

type Identity struct {
	ID       types.Id
	Email    string
	Password string
}

func NewIdentity(id types.Id, email, password string) *Identity {
	return &Identity{
		ID:       id,
		Email:    email,
		Password: password,
	}
}

func NewIdentityFromDto(email, password string) *Identity {
	return &Identity{
		Email:    email,
		Password: password,
	}
}

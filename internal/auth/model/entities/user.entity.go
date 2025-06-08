package entities

import "github.com/GanderBite/reservation-api/internal/pkg/types"

type User struct {
	ID             types.Id
	Name           string
	ReservationIds []*types.Id
	IdentityId     types.Id
}

func NewUser(id types.Id, name string, reservationIds []*types.Id, identityId types.Id) *User {
	return &User{
		ID:             id,
		Name:           name,
		ReservationIds: reservationIds,
		IdentityId:     identityId,
	}
}

func NewUserFromDto(name string) *User {
	return &User{
		Name: name,
	}
}

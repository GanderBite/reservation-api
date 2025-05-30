package domain

import (
	"time"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type Reservation struct {
	ID        types.Id
	Status    ReservationStatus
	Price     types.Price
	CreatedAt time.Time
	ExpiresAt time.Time
	UpdatedAt time.Time
}

func NewReservation(
	Id types.Id,
	status ReservationStatus,
	price types.Price,
	createdAt time.Time,
	expiresAt time.Time,
	updatedAt time.Time,
) *Reservation {
	return &Reservation{
		Id,
		status,
		price,
		createdAt,
		expiresAt,
		updatedAt,
	}
}

func NewReservationFromDto(
	price types.Price,
) *Reservation {
	createdAt := time.Now()
	expiresAt := createdAt.Add(1 * time.Hour)
	return &Reservation{
		Status:    StatusPending,
		Price:     price,
		CreatedAt: createdAt,
		ExpiresAt: expiresAt,
	}
}

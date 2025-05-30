package domain

import (
	"time"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type Reservation struct {
	ID                    types.Id
	Status                ReservationStatus
	Price                 types.Price
	CreatedAt             time.Time
	ExpiresAt             time.Time
	UpdatedAt             time.Time
	AppliedDiscountCodeId *types.Id
}

func NewReservation(
	Id types.Id,
	status ReservationStatus,
	price types.Price,
	createdAt time.Time,
	expiresAt time.Time,
	updatedAt time.Time,
	appliedDiscountCodeId *types.Id,
) *Reservation {
	return &Reservation{
		Id,
		status,
		price,
		createdAt,
		expiresAt,
		updatedAt,
		appliedDiscountCodeId,
	}
}

func (r *Reservation) ApplyDiscount(appliedDiscountId types.Id, amount types.Price) {
	r.AppliedDiscountCodeId = &appliedDiscountId
	r.Price -= amount
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

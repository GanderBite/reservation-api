package dtos

import (
	"time"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

type SeatDto struct {
	ID    types.Id `json:"id"`
	Label string   `json:"label"`
}

type ReservationDto struct {
	ID        types.Id                 `json:"id"`
	Status    domain.ReservationStatus `json:"status"`
	Price     types.Price              `json:"price"`
	CreatedAt time.Time                `json:"createdAt"`
	ExpiresAt time.Time                `json:"expiresAt"`
	Seats     []*SeatDto               `json:"seats"`
}

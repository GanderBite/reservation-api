package domain

import (
	"time"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type ReservedSeat struct {
	ID            types.Id
	ReservationId types.Id
	SeatId        types.Id
	CreatedAt     time.Time
}

func NewReservedSeat(id types.Id, reservationId types.Id, seatId types.Id, createdAt time.Time) *ReservedSeat {
	return &ReservedSeat{
		ID:            id,
		ReservationId: reservationId,
		SeatId:        seatId,
		CreatedAt:     createdAt,
	}
}

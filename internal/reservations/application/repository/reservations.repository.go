package repository

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

type ReservationsRepository interface {
	Insert(ctx context.Context, reservation *domain.Reservation, seatIds []*types.Id) (types.Id, error)
	GetById(ctx context.Context, id types.Id) *domain.Reservation
	GetReservedSeatsByIds(ctx context.Context, ids []*types.Id) ([]*domain.ReservedSeat, error)
}

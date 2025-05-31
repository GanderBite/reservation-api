package repository

import (
	"context"
	"time"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/application/dtos"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

type ReservationsRepository interface {
	Insert(ctx context.Context, reservation *domain.Reservation, seatIds []*types.Id) (types.Id, error)
	GetById(ctx context.Context, id types.Id) (*domain.Reservation, error)
	GetReservationDetails(ctx context.Context, reservationId types.Id) (*dtos.ReservationDto, error)
	GetReservedSeatsByIds(ctx context.Context, ids []*types.Id) ([]*domain.ReservedSeat, error)
	UpdateStatus(ctx context.Context, reservationId types.Id, status domain.ReservationStatus) error
	DeletePending(ctx context.Context, cutoff time.Time) error
	DeleteExpired(ctx context.Context) error
	IsSeatReserved(ctx context.Context, seatId types.Id) (bool, error)
}

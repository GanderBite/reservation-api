package ohs

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
)

type SeatsOHS interface {
	GetAllSeats(ctx context.Context) ([]*entities.Seat, error)
}

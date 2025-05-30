package ohs

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
)

type SeatsOHS interface {
	GetSeatsByIds(ctx context.Context, ids []*types.Id) ([]*entities.Seat, error)
}

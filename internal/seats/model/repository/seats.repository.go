package repository

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
)

type SeatsRepository interface {
	Insert(ctx context.Context, seat *entities.Seat) (types.Id, error)
	GetByIds(ctx context.Context, ids []*types.Id) ([]*entities.Seat, error)
	GetAll(ctx context.Context) ([]*entities.Seat, error)
}

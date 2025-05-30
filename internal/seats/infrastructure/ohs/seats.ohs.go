package ohs

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
	"github.com/GanderBite/reservation-api/internal/seats/model/repository"
)

type SeatsOHS struct {
	repo repository.SeatsRepository
}

func NewSeatsOHS(repo repository.SeatsRepository) *SeatsOHS {
	return &SeatsOHS{repo: repo}
}

func (ohs *SeatsOHS) GetSeatsByIds(ctx context.Context, ids []*types.Id) ([]*entities.Seat, error) {
	return ohs.repo.GetByIds(ctx, ids)
}

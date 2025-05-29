package repository

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
)

type SeatsRepository interface {
	Insert(seat *entities.Seat) (types.Id, error)
	GetAll() ([]*entities.Seat, error)
}

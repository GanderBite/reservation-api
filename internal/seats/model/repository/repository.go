package repository

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/dtos"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
)

type SeatsRepository interface {
	Insert(dto *dtos.CreateSeatDto) (types.Id, error)
	GetAll() ([]*entities.Seat, error)
}

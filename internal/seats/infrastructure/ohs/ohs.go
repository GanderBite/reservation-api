package ohs

import (
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
	"github.com/GanderBite/reservation-api/internal/seats/model/repository"
)

type SeatsOHS struct {
	repo repository.SeatsRepository
}

func NewSeatsOHS(repo repository.SeatsRepository) *SeatsOHS {
	return &SeatsOHS{repo: repo}
}

func (ohs *SeatsOHS) GetAllSeats() ([]*entities.Seat, error) {
	return ohs.repo.GetAll()
}

package ohs

import "github.com/GanderBite/reservation-api/internal/seats/model/entities"

type SeatsOHS interface {
	GetAllSeats() ([]*entities.Seat, error)
}

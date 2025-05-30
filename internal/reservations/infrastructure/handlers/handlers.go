package handlers

import (
	reservation "github.com/GanderBite/reservation-api/internal/reservations/application/repository"
	usecases "github.com/GanderBite/reservation-api/internal/reservations/application/use-cases"
	seats "github.com/GanderBite/reservation-api/internal/seats/model/ohs"
)

type ReservationHandlers struct {
	repo                     reservation.ReservationsRepository
	CreateReservationHandler *createReservationHandler
}

func NewReservationHandlers(reservationRepository reservation.ReservationsRepository, seatsApi seats.SeatsOHS) *ReservationHandlers {
	createReservationUC := usecases.NewCreateReservationUseCase(reservationRepository, seatsApi)
	createReservationHandler := newCreateReservationHandler(createReservationUC)

	return &ReservationHandlers{
		repo:                     reservationRepository,
		CreateReservationHandler: createReservationHandler,
	}
}

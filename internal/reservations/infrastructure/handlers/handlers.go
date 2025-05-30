package handlers

import (
	discount_codes "github.com/GanderBite/reservation-api/internal/discount-codes/model/ohs"
	reservation "github.com/GanderBite/reservation-api/internal/reservations/application/repository"
	usecases "github.com/GanderBite/reservation-api/internal/reservations/application/use-cases"
	"github.com/GanderBite/reservation-api/internal/reservations/infrastructure/services"
	seats "github.com/GanderBite/reservation-api/internal/seats/model/ohs"
)

type ReservationHandlers struct {
	repo                     reservation.ReservationsRepository
	CreateReservationHandler *createReservationHandler
}

func NewReservationHandlers(
	reservationRepository reservation.ReservationsRepository,
	seatsApi seats.SeatsOHS,
	discountCodesApi discount_codes.DiscountCodeOhs,
) *ReservationHandlers {
	createReservationUC := usecases.NewCreateReservationUseCase(
		reservationRepository,
		seatsApi,
		discountCodesApi,
		services.NewApplyDiscountService(),
	)
	createReservationHandler := newCreateReservationHandler(createReservationUC)

	return &ReservationHandlers{
		repo:                     reservationRepository,
		CreateReservationHandler: createReservationHandler,
	}
}

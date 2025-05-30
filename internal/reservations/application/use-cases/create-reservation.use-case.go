package usecases

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/application/dtos"
	"github.com/GanderBite/reservation-api/internal/reservations/application/repository"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
	"github.com/GanderBite/reservation-api/internal/seats/model/ohs"
	"github.com/google/uuid"
)

type CreateReservationUseCase struct {
	reservationRepository repository.ReservationsRepository
	seatsRepository       ohs.SeatsOHS
}

func NewCreateReservationUseCase(reservationRepository repository.ReservationsRepository, seatsRepository ohs.SeatsOHS) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		reservationRepository,
		seatsRepository,
	}
}

func (uc *CreateReservationUseCase) Execute(ctx context.Context, dto *dtos.CreateReservationDto) (types.Id, error) {
	seats, err := uc.seatsRepository.GetSeatsByIds(ctx, dto.SeatIds)
	if err != nil {
		return uuid.Nil, err
	}

	reservedSeats, err := uc.reservationRepository.GetReservedSeatsByIds(ctx, dto.SeatIds)
	if err != nil {
		return uuid.Nil, err
	}

	if len(reservedSeats) > 0 {
		return uuid.Nil, domain.ErrSeatsAlreadyReserved
	}

	var totalPrice types.Price = 0.0
	for _, seat := range seats {
		totalPrice += seat.Price
	}

	reservation := domain.NewReservationFromDto(totalPrice)
	reservationId, err := uc.reservationRepository.Insert(ctx, reservation, dto.SeatIds)
	if err != nil {
		return uuid.Nil, err
	}

	return reservationId, nil
}

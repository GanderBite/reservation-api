package usecases

import (
	"context"

	"github.com/google/uuid"

	discountCodesOhs "github.com/GanderBite/reservation-api/internal/discount-codes/model/ohs"
	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/application/dtos"
	"github.com/GanderBite/reservation-api/internal/reservations/application/repository"
	"github.com/GanderBite/reservation-api/internal/reservations/application/services"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
	"github.com/GanderBite/reservation-api/internal/seats/model/ohs"
)

type CreateReservationUseCase struct {
	reservationRepository   repository.ReservationsRepository
	seatsRepository         ohs.SeatsOHS
	discountCodesRepository discountCodesOhs.DiscountCodeOhs
	applyDiscountService    services.ApplyDiscountService
}

func NewCreateReservationUseCase(
	reservationRepository repository.ReservationsRepository,
	seatsRepository ohs.SeatsOHS,
	discountCodesRepository discountCodesOhs.DiscountCodeOhs,
	applyDiscountService services.ApplyDiscountService,
) *CreateReservationUseCase {
	return &CreateReservationUseCase{
		reservationRepository,
		seatsRepository,
		discountCodesRepository,
		applyDiscountService,
	}
}

func (uc *CreateReservationUseCase) Execute(ctx context.Context, dto *dtos.CreateReservationDto) (types.Id, error) {
	if len(dto.SeatIds) == 0 {
		return uuid.Nil, domain.ErrMissingSeats
	}

	seats, err := uc.seatsRepository.GetSeatsByIds(ctx, dto.SeatIds)
	if err != nil {
		return uuid.Nil, err
	}

	seatsCount := len(seats)
	if seatsCount == 0 {
		return uuid.Nil, domain.ErrMissingSeats
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

	discountCode := uc.applyDiscountService.ApplyDiscountBySeatsCount(len(seats))
	if discountCode != "" {
		discount, err := uc.discountCodesRepository.GetDiscountCodeByCode(ctx, discountCode)
		if err != nil {
			return uuid.Nil, err
		}

		if discount != nil {
			reservation.ApplyDiscount(discount.ID, discount.Price)
		}
	}

	reservationId, err := uc.reservationRepository.Insert(ctx, reservation, dto.SeatIds)
	if err != nil {
		return uuid.Nil, err
	}

	return reservationId, nil
}

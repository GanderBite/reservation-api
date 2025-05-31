package usecases

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/application/repository"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

type UpdateReservationStatusUseCase struct {
	repo repository.ReservationsRepository
}

func NewUpdateReservationStatusUseCase(
	repo repository.ReservationsRepository,
) *UpdateReservationStatusUseCase {
	return &UpdateReservationStatusUseCase{
		repo,
	}
}

func (uc *UpdateReservationStatusUseCase) Execute(
	ctx context.Context,
	reservationId types.Id,
	status domain.ReservationStatus,
) error {
	reservation, err := uc.repo.GetById(ctx, reservationId)
	if err != nil {
		return err
	}

	if reservation == nil {
		return domain.ErrReservationNotFound
	}

	err = reservation.UpdateStatus(status)
	if err != nil {
		return err
	}

	err = uc.repo.UpdateStatus(ctx, reservation.ID, reservation.Status)

	return err
}

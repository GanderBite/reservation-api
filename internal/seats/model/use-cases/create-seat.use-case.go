package usecase

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/dtos"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
	"github.com/GanderBite/reservation-api/internal/seats/model/repository"
)

type CreateSeatUseCase struct {
	seatsRepository repository.SeatsRepository
}

func NewCreateSeatUseCase(seatsRepository repository.SeatsRepository) *CreateSeatUseCase {
	return &CreateSeatUseCase{
		seatsRepository: seatsRepository,
	}
}

func (uc *CreateSeatUseCase) Execute(ctx context.Context, dto *dtos.CreateSeatDto) (types.Id, error) {
	seat := entities.NewSeatFromDto(dto.Row, dto.Col, dto.Price)
	seatId, err := uc.seatsRepository.Insert(ctx, seat)
	if err != nil {
		return types.Id{}, err
	}

	return seatId, nil
}

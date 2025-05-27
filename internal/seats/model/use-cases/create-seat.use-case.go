package usecase

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/dtos"
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

func (uc *CreateSeatUseCase) Execute(dto *dtos.CreateSeatDto) (types.Id, error) {
	return uc.seatsRepository.Insert(dto)
}

package handlers

import (
	"github.com/GanderBite/reservation-api/internal/seats/model/repository"
	usecase "github.com/GanderBite/reservation-api/internal/seats/model/use-cases"
)

type SeatHandlers struct {
	repo              repository.SeatsRepository
	CreateSeatHandler *createSeatHandler
}

func NewSeatHandlers(repo repository.SeatsRepository) *SeatHandlers {
	createSeatUseCase := usecase.NewCreateSeatUseCase(repo)
	createSeatHandler := newCreateSeatHandler(createSeatUseCase)

	return &SeatHandlers{
		repo:              repo,
		CreateSeatHandler: createSeatHandler,
	}
}

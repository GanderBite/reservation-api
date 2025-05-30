package services

import (
	"context"
	"log"
	"time"

	"github.com/GanderBite/reservation-api/internal/reservations/application/repository"
)

type CleanUpService struct {
	reservationsRepo repository.ReservationsRepository
}

func NewCleanUpService(reservationsRepo repository.ReservationsRepository) *CleanUpService {
	return &CleanUpService{
		reservationsRepo: reservationsRepo,
	}
}

func (s *CleanUpService) DeleteLingeringReservations() {
	ctx := context.Background()
	log.Println("Cleaning up pedning reservations")

	if err := s.reservationsRepo.DeleteExpired(ctx); err != nil {
		log.Printf("Error deleting expired reservations: %v", err)
	}

	cutoff := time.Now().Add(-15 * time.Minute)
	if err := s.reservationsRepo.DeletePending(ctx, cutoff); err != nil {
		log.Printf("Error deleting old pending reservations: %v", err)
	}
}

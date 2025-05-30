package reservations

import (
	"database/sql"

	discountCodesOhs "github.com/GanderBite/reservation-api/internal/discount-codes/model/ohs"
	"github.com/GanderBite/reservation-api/internal/reservations/application/repository"
	"github.com/GanderBite/reservation-api/internal/reservations/infrastructure/database"
	"github.com/GanderBite/reservation-api/internal/reservations/infrastructure/handlers"
	"github.com/GanderBite/reservation-api/internal/seats/model/ohs"
)

type ReservationsModule struct {
	repo     repository.ReservationsRepository
	Handlers *handlers.ReservationHandlers
}

func NewReservationsModule(db *sql.DB, seatsApi ohs.SeatsOHS, discountCodesOhs discountCodesOhs.DiscountCodeOhs) *ReservationsModule {
	repo := database.NewPostgresReservationsRepository(db)
	handlers := handlers.NewReservationHandlers(repo, seatsApi, discountCodesOhs)

	return &ReservationsModule{
		repo,
		handlers,
	}
}

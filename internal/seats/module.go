package seats

import (
	"database/sql"

	"github.com/GanderBite/reservation-api/internal/seats/infrastructure/database"
	"github.com/GanderBite/reservation-api/internal/seats/infrastructure/handlers"
	"github.com/GanderBite/reservation-api/internal/seats/infrastructure/ohs"
	ohsModel "github.com/GanderBite/reservation-api/internal/seats/model/ohs"
	"github.com/GanderBite/reservation-api/internal/seats/model/repository"
)

type SeatsModule struct {
	repo     repository.SeatsRepository
	Api      ohsModel.SeatsOHS
	Handlers *handlers.SeatHandlers
}

func NewSeatsModule(db *sql.DB) *SeatsModule {
	repo := database.NewPostgresSeatsRepository(db)
	Api := ohs.NewSeatsOHS(repo)
	handlers := handlers.NewSeatHandlers(repo)

	return &SeatsModule{
		repo,
		Api,
		handlers,
	}
}

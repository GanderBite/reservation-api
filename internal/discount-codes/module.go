package discount_code

import (
	"database/sql"

	"github.com/GanderBite/reservation-api/internal/discount-codes/infrastructure/database"
	"github.com/GanderBite/reservation-api/internal/discount-codes/infrastructure/handlers"
	ohsInfra "github.com/GanderBite/reservation-api/internal/discount-codes/infrastructure/ohs"
	"github.com/GanderBite/reservation-api/internal/discount-codes/model/ohs"
	"github.com/GanderBite/reservation-api/internal/discount-codes/model/repository"
)

type DiscountCodesModule struct {
	repo     repository.DiscountCodesRepository
	Handlers *handlers.DiscountCodeHandlers
	Api      ohs.DiscountCodeOhs
}

func NewDiscountModule(db *sql.DB) *DiscountCodesModule {
	repo := database.NewPostgressDiscountCodesRepository(db)
	api := ohsInfra.NewDiscountCodesOhs(repo)
	handlers := handlers.NewDiscountCodeHandlers(repo)

	return &DiscountCodesModule{
		repo:     repo,
		Handlers: handlers,
		Api:      api,
	}
}

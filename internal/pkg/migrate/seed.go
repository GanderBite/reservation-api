package migrate

import (
	"database/sql"
	"log"
	"time"

	"github.com/GanderBite/reservation-api/internal/discount-codes/model/entities"
	"github.com/GanderBite/reservation-api/internal/pkg/env"
	seatEntities "github.com/GanderBite/reservation-api/internal/seats/model/entities"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func getSeedSeats() []*seatEntities.Seat {
	return []*seatEntities.Seat{
		seatEntities.NewSeat(uuid.New(), "A", 1, 15),
		seatEntities.NewSeat(uuid.New(), "A", 2, 20),
		seatEntities.NewSeat(uuid.New(), "A", 3, 20),
		seatEntities.NewSeat(uuid.New(), "A", 4, 20),
		seatEntities.NewSeat(uuid.New(), "A", 5, 15),
		seatEntities.NewSeat(uuid.New(), "B", 1, 10),
		seatEntities.NewSeat(uuid.New(), "B", 2, 15),
		seatEntities.NewSeat(uuid.New(), "B", 3, 15),
		seatEntities.NewSeat(uuid.New(), "B", 4, 15),
		seatEntities.NewSeat(uuid.New(), "B", 5, 10),
	}
}

func getSeedDiscountCodes() []*entities.DiscountCode {
	return []*entities.DiscountCode{
		entities.NewDiscountCode(uuid.New(), "PAIR", time.Now(), 5),
		entities.NewDiscountCode(uuid.New(), "BIG_FAM", time.Now(), 10),
	}
}

func Seed() {
	db, err := sql.Open("pgx", env.GetEnvString("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	seats := getSeedSeats()

	seatsQuery := `
		INSERT INTO seats (id, row, col, price)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (row, col) DO NOTHING
	`

	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	for _, s := range seats {
		_, err := tx.Exec(seatsQuery,
			s.ID,
			s.Row,
			s.Col,
			s.Price,
		)
		if err != nil {
			log.Fatalf("failed to upsert seat %+v: %v", s, err)
		}
	}

	discountCodes := getSeedDiscountCodes()

	discountCodesQuery := `
		INSERT INTO discount_codes (id, code, price, created_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (code) DO NOTHING
	`

	for _, dc := range discountCodes {
		_, err := tx.Exec(discountCodesQuery,
			dc.ID,
			dc.Code,
			dc.Price,
			dc.CreatedAt,
		)
		if err != nil {
			log.Fatalf("failed to upsert discount codes %+v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("failed to commit transaction: %v", err)
	}

	log.Println("Seed data inserted/upserted successfully.")

}

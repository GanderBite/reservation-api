package migrate

import (
	"database/sql"
	"log"

	"github.com/GanderBite/reservation-api/internal/pkg/env"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
	"github.com/google/uuid"
)

func getSeedSeats() []*entities.Seat {
	return []*entities.Seat{
		entities.NewSeat(uuid.New(), "A", 1, 15),
		entities.NewSeat(uuid.New(), "A", 2, 20),
		entities.NewSeat(uuid.New(), "A", 3, 20),
		entities.NewSeat(uuid.New(), "A", 4, 20),
		entities.NewSeat(uuid.New(), "A", 5, 15),
		entities.NewSeat(uuid.New(), "B", 1, 10),
		entities.NewSeat(uuid.New(), "B", 2, 15),
		entities.NewSeat(uuid.New(), "B", 3, 15),
		entities.NewSeat(uuid.New(), "B", 4, 15),
		entities.NewSeat(uuid.New(), "B", 5, 10),
	}
}

func Seed() {
	db, err := sql.Open("postgres", env.GetEnvString("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	seats := getSeedSeats()

	seatsQuery := `
		INSERT INTO seats (id, row, col, price)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO NOTHING
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

	if err := tx.Commit(); err != nil {
		log.Fatalf("failed to commit transaction: %v", err)
	}

	log.Println("Seed data inserted/upserted successfully.")

}

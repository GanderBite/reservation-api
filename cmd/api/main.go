package main

import (
	"database/sql"
	"log"

	discount_code "github.com/GanderBite/reservation-api/internal/discount-codes"
	"github.com/GanderBite/reservation-api/internal/pkg/env"
	"github.com/GanderBite/reservation-api/internal/reservations"
	"github.com/GanderBite/reservation-api/internal/seats"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	port          int
	seats         *seats.SeatsModule
	reservations  *reservations.ReservationsModule
	discountCodes *discount_code.DiscountCodesModule
}

func main() {
	env.Load()

	db, err := sql.Open("pgx", env.GetEnvString("DATABASE_URL"))
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return
	}

	seats := seats.NewSeatsModule(db)
	reservations := reservations.NewReservationsModule(db, seats.Api)
	discountCodes := discount_code.NewDiscountModule(db)

	app := &application{
		port:          env.GetEnvInt("PORT"),
		seats:         seats,
		reservations:  reservations,
		discountCodes: discountCodes,
	}

	err = app.serve()
	if err != nil {
		log.Println("Failed to start a server:", err)
	}

}

package main

import (
	"database/sql"
	"log"

	discount_code "github.com/GanderBite/reservation-api/internal/discount-codes"
	"github.com/GanderBite/reservation-api/internal/pkg/env"
	"github.com/GanderBite/reservation-api/internal/pkg/services"
	"github.com/GanderBite/reservation-api/internal/reservations"
	"github.com/GanderBite/reservation-api/internal/seats"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/robfig/cron/v3"
)

// @title       Reservation API
// @version     1.0
// @description API for managing reservations
// @BasePath    /api/v1

type application struct {
	port          int
	seats         *seats.SeatsModule
	reservations  *reservations.ReservationsModule
	discountCodes *discount_code.DiscountCodesModule
}

func main() {
	env.Load()
	appEnv := env.GetEnvString("APP_ENV")

	db, err := sql.Open("pgx", env.GetEnvString("DATABASE_URL"))
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return
	}

	seats := seats.NewSeatsModule(db)
	discountCodes := discount_code.NewDiscountModule(db)
	reservations := reservations.NewReservationsModule(db, seats.Api, discountCodes.Api)
	cleanUpService := services.NewCleanUpService(reservations.Repo)

	app := &application{
		port:          env.GetEnvInt("PORT"),
		seats:         seats,
		reservations:  reservations,
		discountCodes: discountCodes,
	}

	c := cron.New()

	spec := "*/15 * * * *"
	if appEnv == "dev" {
		spec = "* * * * *"
	}

	_, err = c.AddFunc(spec, cleanUpService.DeleteLingeringReservations)
	if err != nil {
		panic(err)
	}

	c.Start()

	err = app.serve()
	if err != nil {
		log.Println("Failed to start a server:", err)
	}
}

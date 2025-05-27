package main

import (
	"database/sql"
	"log"

	"github.com/GanderBite/reservation-api/internal/pkg/env"
	"github.com/GanderBite/reservation-api/internal/seats"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	port  int
	seats *seats.SeatsModule
}

func main() {
	env.Load()

	db, err := sql.Open("pgx", env.GetEnvString("DATABASE_URL"))
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return
	}

	seats := seats.NewSeatsModule(db)

	app := &application{
		port:  env.GetEnvInt("PORT"),
		seats: seats,
	}

	err = app.serve()
	if err != nil {
		log.Println("Failed to start a server:", err)
	}

}

package main

import (
	"log"
	"os"

	"github.com/GanderBite/reservation-api/internal/pkg/env"
	"github.com/GanderBite/reservation-api/internal/pkg/migrate"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a command: up | down | create <name>")
	}

	command := os.Args[1]

	env.Load()

	switch command {
	case "up", "down":
		migrate.RunMigrations(command)
	case "create":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a migration name: create <name>")
		}
		migrate.CreateMigration(os.Args[2])
	case "seed":
		migrate.Seed()
	default:
		log.Fatalf("Unknown command: %s", command)
	}
}

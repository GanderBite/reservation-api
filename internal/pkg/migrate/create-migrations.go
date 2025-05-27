package migrate

import (
	"log"
	"os"
	"os/exec"
)

func CreateMigration(name string) {
	cmd := exec.Command(
		"migrate",
		"create",
		"-ext", "sql",
		"-dir", "./cmd/migrate/migrations",
		"-seq", name,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to create migration: %v", err)
	}
}

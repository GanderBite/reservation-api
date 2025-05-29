package database

import (
	"context"
	"database/sql"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/seats/model/entities"
	"github.com/google/uuid"
)

type PostgresSeatsRepository struct {
	db *sql.DB
}

func NewPostgresSeatsRepository(db *sql.DB) *PostgresSeatsRepository {
	return &PostgresSeatsRepository{
		db,
	}
}

func (repo *PostgresSeatsRepository) Insert(ctx context.Context, seat *entities.Seat) (types.Id, error) {
	query := `INSERT INTO seats
		(row, col, price)
	VALUES
		($1, $2, $3)
	RETURNING id`

	var createdSeatId string
	err := repo.db.QueryRowContext(ctx, query, seat.Row, seat.Col, seat.Price).Scan(&createdSeatId)
	if err != nil {
		return uuid.Nil, err
	}

	parsedId, err := uuid.Parse(createdSeatId)
	if err != nil {
		return uuid.Nil, err
	}

	return parsedId, nil
}

func (repo *PostgresSeatsRepository) GetAll(ctx context.Context) ([]*entities.Seat, error) {
	query := `SELECT id, row, col, price FROM seats`

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []*entities.Seat

	for rows.Next() {
		var id uuid.UUID
		var row string
		var col int
		var price types.Price

		err := rows.Scan(&id, &row, &col, &price)
		if err != nil {
			return nil, err
		}

		seat := entities.NewSeat(id, row, col, price)
		seats = append(seats, seat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seats, nil
}

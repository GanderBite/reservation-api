package database

import (
	"context"
	"database/sql"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
	"github.com/google/uuid"
)

type PostgresReservationsRepository struct {
	db *sql.DB
}

func NewPostgresReservationsRepository(db *sql.DB) *PostgresReservationsRepository {
	return &PostgresReservationsRepository{db}
}

func (repo *PostgresReservationsRepository) Insert(ctx context.Context, r *domain.Reservation, seatIds []*types.Id) (types.Id, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback()

	insertReservation := `
		INSERT INTO reservations (status, price, created_at, expires_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	var reservationId string
	err = tx.QueryRowContext(ctx, insertReservation, r.Status, r.Price, r.CreatedAt, r.ExpiresAt).Scan(&reservationId)
	if err != nil {
		return uuid.Nil, err
	}

	parsedId, err := uuid.Parse(reservationId)
	if err != nil {
		return uuid.Nil, err
	}

	insertReservedSeats := `
		INSERT INTO reserved_seats (reservation_id, seat_id, created_at)
		VALUES ($1, $2, NOW())`

	for _, seatId := range seatIds {
		_, err := tx.ExecContext(ctx, insertReservedSeats, parsedId, *seatId)
		if err != nil {
			return uuid.Nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return uuid.Nil, err
	}

	return parsedId, nil
}

func (repo *PostgresReservationsRepository) GetById(ctx context.Context, id types.Id) *domain.Reservation {
	return nil
}

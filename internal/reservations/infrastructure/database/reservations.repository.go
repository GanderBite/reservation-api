package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/application/dtos"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

type PostgresReservationsRepository struct {
	db *sql.DB
}

func NewPostgresReservationsRepository(db *sql.DB) *PostgresReservationsRepository {
	return &PostgresReservationsRepository{db}
}

func (repo *PostgresReservationsRepository) Insert(
	ctx context.Context,
	r *domain.Reservation,
	seatIds []*types.Id,
) (types.Id, error) {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	defer func() {
		if cerr := tx.Rollback(); cerr != nil {
			log.Fatalln(cerr.Error())
		}
	}()

	insertReservation := `
		INSERT INTO reservations (status, price, created_at, expires_at, applied_discount_code_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`
	var appliedDiscountCode any
	if r.AppliedDiscountCodeId != nil {
		appliedDiscountCode = *r.AppliedDiscountCodeId
	} else {
		appliedDiscountCode = nil
	}
	var reservationId string
	err = tx.QueryRowContext(ctx, insertReservation, r.Status, r.Price, r.CreatedAt, r.ExpiresAt, appliedDiscountCode).
		Scan(&reservationId)
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

func (repo *PostgresReservationsRepository) GetReservedSeatsByIds(
	ctx context.Context,
	ids []*types.Id,
) ([]*domain.ReservedSeat, error) {
	if len(ids) == 0 {
		return []*domain.ReservedSeat{}, nil
	}

	placeholders := make([]string, len(ids))
	args := make([]any, len(ids))
	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT id, reservation_id, seat_id, created_at
		FROM reserved_seats
		WHERE seat_id IN (%s)`, strings.Join(placeholders, ", "))

	rows, err := repo.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatalln(cerr.Error())
		}
	}()

	var rsArr []*domain.ReservedSeat
	for rows.Next() {
		var id types.Id
		var reservationId types.Id
		var seatId types.Id
		var createdAt time.Time

		if err := rows.Scan(&id, &reservationId, &seatId, &createdAt); err != nil {
			return nil, err
		}

		rs := domain.NewReservedSeat(id, reservationId, seatId, createdAt)
		rsArr = append(rsArr, rs)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rsArr, nil
}

func (repo *PostgresReservationsRepository) GetById(ctx context.Context, id types.Id) (*domain.Reservation, error) {
	query := `
		SELECT id, status, price, created_at, applied_discount_code_id
		FROM reservations
		WHERE id = $1
	`

	var (
		reservationId         uuid.UUID
		status                domain.ReservationStatus
		price                 types.Price
		createdAt             time.Time
		appliedDiscountCodeId *uuid.UUID
	)

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&reservationId,
		&status,
		&price,
		&createdAt,
		&appliedDiscountCodeId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	res := &domain.Reservation{
		ID:                    types.Id(reservationId),
		Status:                status,
		Price:                 price,
		CreatedAt:             createdAt,
		AppliedDiscountCodeId: nil,
	}

	if appliedDiscountCodeId != nil {
		id := types.Id(*appliedDiscountCodeId)
		res.AppliedDiscountCodeId = &id
	}

	return res, nil
}

func (repo *PostgresReservationsRepository) UpdateStatus(
	ctx context.Context,
	reservationId types.Id,
	status domain.ReservationStatus,
) error {

	confirmQuery := `
			UPDATE reservations
			SET status = $1
			WHERE id = $2
		`
	deleteQuery := `
		DELETE FROM reservations
		WHERE id = $1;
	`

	var err error
	switch status {
	case domain.StatusExpired:
		_, err = repo.db.ExecContext(ctx, deleteQuery, reservationId)
	default:
		_, err = repo.db.ExecContext(ctx, confirmQuery, status, reservationId)
	}

	return err
}

func (repo *PostgresReservationsRepository) DeletePending(ctx context.Context, cutoff time.Time) error {
	_, err := repo.db.ExecContext(ctx, `
		DELETE FROM reservations
		WHERE status = 'pending' AND created_at < $1
	`, cutoff)
	return err
}

func (repo *PostgresReservationsRepository) DeleteExpired(ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, `
		DELETE FROM reservations
		WHERE status = 'expired' OR expires_at <= NOW()
	`)
	return err
}

func (repo *PostgresReservationsRepository) GetReservationDetails(
	ctx context.Context,
	reservationId types.Id,
) (*dtos.ReservationDto, error) {
	query := `
		SELECT
			r.id,
			r.status,
			r.price,
			r.created_at,
			r.expires_at,
			s.id AS seat_id,
			s.row,
			s.col,
			s.price
		FROM reservations r
		JOIN reserved_seats rs ON rs.reservation_id = r.id
		JOIN seats s ON s.id = rs.seat_id
		WHERE r.id = $1;
	`

	rows, err := repo.db.QueryContext(ctx, query, reservationId)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatalln(cerr.Error())
		}
	}()

	var (
		resDto *dtos.ReservationDto
		seats  []*dtos.SeatDto
	)

	for rows.Next() {
		var (
			id        uuid.UUID
			status    string
			price     types.Price
			createdAt time.Time
			expiresAt time.Time

			seatId    uuid.UUID
			row       string
			col       int
			seatPrice types.Price
		)

		err := rows.Scan(&id, &status, &price, &createdAt, &expiresAt, &seatId, &row, &col, &seatPrice)
		if err != nil {
			return nil, err
		}

		// Only initialize once
		if resDto == nil {
			resDto = &dtos.ReservationDto{
				ID:        id,
				Status:    domain.ReservationStatus(status),
				Price:     price,
				CreatedAt: createdAt,
				ExpiresAt: expiresAt,
			}
		}

		seats = append(seats, &dtos.SeatDto{
			ID:    seatId,
			Label: fmt.Sprintf("%s%d", row, col),
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if resDto == nil {
		return nil, domain.ErrReservationNotFound
	}

	resDto.Seats = seats
	return resDto, nil
}

func (repo *PostgresReservationsRepository) IsSeatReserved(ctx context.Context, seatId types.Id) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM reserved_seats rs
			JOIN reservations r ON rs.reservation_id = r.id
			WHERE rs.seat_id = $1
			AND r.status != 'expired'
		);
	`

	var exists bool
	err := repo.db.QueryRowContext(ctx, query, seatId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

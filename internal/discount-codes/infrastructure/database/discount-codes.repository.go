package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/GanderBite/reservation-api/internal/discount-codes/model/entities"
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type PostgresDiscountCodesRepository struct {
	db *sql.DB
}

func NewPostgressDiscountCodesRepository(db *sql.DB) *PostgresDiscountCodesRepository {
	return &PostgresDiscountCodesRepository{db}
}

func (repo *PostgresDiscountCodesRepository) GetByCode(
	ctx context.Context,
	code string,
) (*entities.DiscountCode, error) {
	query := `SELECT id, code, price, created_at FROM discount_codes WHERE code = $1`

	row := repo.db.QueryRowContext(ctx, query, code)

	var id types.Id
	var discountCode string
	var createdAt time.Time
	var price types.Price

	err := row.Scan(&id, &discountCode, &price, &createdAt)
	if err != nil {
		return nil, err
	}

	return entities.NewDiscountCode(id, discountCode, createdAt, price), nil
}

func (repo *PostgresDiscountCodesRepository) GetAll(ctx context.Context) ([]*entities.DiscountCode, error) {
	query := `SELECT id, code, price, created_at FROM discount_codes`

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Fatalln(cerr.Error())
		}
	}()

	var discountCodes []*entities.DiscountCode

	for rows.Next() {
		var id uuid.UUID
		var code string
		var price types.Price
		var createdAt time.Time

		if err := rows.Scan(&id, &code, &price, &createdAt); err != nil {
			return nil, err
		}

		discountCode := entities.NewDiscountCode(id, code, createdAt, price)
		discountCodes = append(discountCodes, discountCode)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return discountCodes, nil
}

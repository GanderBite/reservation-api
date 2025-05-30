package repository

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/discount-codes/model/entities"
)

type DiscountCodesRepository interface {
	GetByCode(ctx context.Context, code string) (*entities.DiscountCode, error)
	GetAll(ctx context.Context) ([]*entities.DiscountCode, error)
}

package ohs

import (
	"golang.org/x/net/context"

	"github.com/GanderBite/reservation-api/internal/discount-codes/model/entities"
)

type DiscountCodeOhs interface {
	GetDiscountCodeByCode(ctx context.Context, code string) (*entities.DiscountCode, error)
}

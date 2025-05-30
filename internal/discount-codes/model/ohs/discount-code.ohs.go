package ohs

import (
	"github.com/GanderBite/reservation-api/internal/discount-codes/model/entities"
	"golang.org/x/net/context"
)

type DiscountCodeOhs interface {
	GetDiscountCodeByCode(ctx context.Context, code string) (*entities.DiscountCode, error)
}

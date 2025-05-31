package entities

import (
	"time"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type DiscountCode struct {
	ID        types.Id    `json:"id" example:"discount-code-uuid"`
	Code      string      `json:"code" example:"BIG_FAM"`
	CreatedAt time.Time   `json:"createdAt"`
	Price     types.Price `json:"price" example:"10"`
}

func NewDiscountCode(id types.Id, code string, createdAt time.Time, price types.Price) *DiscountCode {
	return &DiscountCode{
		ID:        id,
		Code:      code,
		CreatedAt: createdAt,
		Price:     price,
	}
}

func NewDiscountCodeFromDto(code string, price types.Price) *DiscountCode {
	return &DiscountCode{
		Code:      code,
		CreatedAt: time.Now(),
		Price:     price,
	}
}

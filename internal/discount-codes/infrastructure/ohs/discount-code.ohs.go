package ohs

import (
	"context"

	"github.com/GanderBite/reservation-api/internal/discount-codes/model/entities"
	"github.com/GanderBite/reservation-api/internal/discount-codes/model/repository"
)

type DiscountCodesOhs struct {
	repo repository.DiscountCodesRepository
}

func NewDiscountCodesOhs(repo repository.DiscountCodesRepository) *DiscountCodesOhs {
	return &DiscountCodesOhs{repo}
}

func (ohs *DiscountCodesOhs) GetDiscountCodeByCode(ctx context.Context, code string) (*entities.DiscountCode, error) {
	discountCode, err := ohs.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if discountCode == nil {
		return nil, entities.ErrNoFoundDiscountCode
	}

	return discountCode, nil
}

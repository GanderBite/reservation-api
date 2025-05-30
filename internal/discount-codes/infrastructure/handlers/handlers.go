package handlers

import "github.com/GanderBite/reservation-api/internal/discount-codes/model/repository"

type DiscountCodeHandlers struct {
	repo repository.DiscountCodesRepository
}

func NewDiscountCodeHandlers(repo repository.DiscountCodesRepository) *DiscountCodeHandlers {
	return &DiscountCodeHandlers{repo}
}

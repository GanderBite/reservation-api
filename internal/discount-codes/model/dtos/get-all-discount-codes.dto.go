package dtos

import "github.com/GanderBite/reservation-api/internal/discount-codes/model/entities"

type GetAllDiscountCodesResponse struct {
	Code   int                      `json:"code" example:"200"`
	Status string                   `json:"status" example:"success"`
	Data   []*entities.DiscountCode `json:"data"`
}

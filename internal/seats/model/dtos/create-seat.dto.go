package dtos

import "github.com/GanderBite/reservation-api/internal/pkg/types"

type CreateSeatDto struct {
	Row   string  `json:"row"   binding:"required,max=1"`
	Col   int     `json:"col"   binding:"required,min=1"`
	Price float64 `json:"price" binding:"required,min=1"`
}

type CreateSeatResponse struct {
	Code   int      `json:"code"   example:"201"`
	Status string   `json:"status" example:"success"`
	Data   types.Id `json:"data"   example:"created-seat-uuid"`
}

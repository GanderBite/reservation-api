package dtos

import "github.com/GanderBite/reservation-api/internal/pkg/types"

type CreateReservationDto struct {
	SeatIds []*types.Id `json:"seatIds" binding:"required,min=1,dive,required,uuid"`
}

type CreateReservationResponse struct {
	Code   int      `json:"code" example:"201"`
	Status string   `json:"status" example:"success"`
	Data   types.Id `json:"data" example:"created-reservation-uuid"`
}

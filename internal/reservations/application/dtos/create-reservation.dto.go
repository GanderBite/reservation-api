package dtos

import "github.com/GanderBite/reservation-api/internal/pkg/types"

type CreateReservationDto struct {
	SeatIds []*types.Id `json:"seatIds" binding:"required,min=1,dive,required,uuid"`
}

package dtos

import "github.com/GanderBite/reservation-api/internal/seats/model/entities"

type GetAllSeatsResponse struct {
	Code   int              `json:"code" example:"200"`
	Status string           `json:"status" example:"success"`
	Data   []*entities.Seat `json:"data"`
}

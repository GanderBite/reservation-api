package dtos

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type ConfirmReservationStatusDto struct {
	ReservationId types.Id `json:"reservationId" binding:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
}

type CancelReservationStatusDto struct {
	ReservationId types.Id `json:"reservationId" binding:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
}

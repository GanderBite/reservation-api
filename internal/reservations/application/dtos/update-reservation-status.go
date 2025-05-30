package dtos

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type ConfirmReservationStatusDto struct {
	ReservationId types.Id `json:"reservationId" binding:"required,uuid"`
}

type CancelReservationStatusDto struct {
	ReservationId types.Id `json:"reservationId" binding:"required,uuid"`
}

package dtos

import (
	"time"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

type SeatDto struct {
	ID    types.Id `json:"id"    example:"550e8400-e29b-41d4-a716-446655440000"`
	Label string   `json:"label" example:"A1"`
}

type ReservationDto struct {
	ID        types.Id                 `json:"id"        example:"550e8400-e29b-41d4-a716-446655440000"`
	Status    domain.ReservationStatus `json:"status"    example:"confirmed"`
	Price     types.Price              `json:"price"     example:"35.00"`
	CreatedAt time.Time                `json:"createdAt"`
	ExpiresAt time.Time                `json:"expiresAt"`
	Seats     []*SeatDto               `json:"seats"`
}

type ReservationSuccessResponse struct {
	Code   int             `json:"code"   example:"200"`
	Status string          `json:"status" example:"success"`
	Data   *ReservationDto `json:"data"`
}

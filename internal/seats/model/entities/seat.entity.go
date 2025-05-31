package entities

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type Seat struct {
	ID    types.Id    `json:"id"    example:"seat-uuid"`
	Row   string      `json:"row"   example:"A"`
	Col   int         `json:"col"   example:"1"`
	Price types.Price `json:"price" example:"15"`
}

func NewSeat(id types.Id, row string, col int, price types.Price) *Seat {
	return &Seat{
		ID:    id,
		Row:   row,
		Col:   col,
		Price: price,
	}
}

func NewSeatFromDto(row string, col int, price float64) *Seat {
	return &Seat{
		Row:   row,
		Col:   col,
		Price: types.Price(price),
	}
}

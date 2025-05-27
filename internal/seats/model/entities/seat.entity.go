package entities

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type Seat struct {
	ID    types.Id
	Row   string
	Col   int
	Price types.Price
}

func NewSeat(id types.Id, row string, col int, price types.Price) *Seat {
	return &Seat{
		id,
		row,
		col,
		price,
	}
}

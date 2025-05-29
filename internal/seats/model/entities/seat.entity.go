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

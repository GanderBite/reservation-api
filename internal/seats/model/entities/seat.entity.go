package entities

import (
	"github.com/GanderBite/reservation-api/internal/pkg/types"
)

type Seat struct {
	ID    types.Id    `json:"id"`
	Row   string      `json:"row"`
	Col   int         `json:"col"`
	Price types.Price `json:"price"`
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

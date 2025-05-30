package domain

import "errors"

var (
	ErrSeatsAlreadyReserved = errors.New("one or more selected seats are already reserved")
	ErrMissingSeats         = errors.New("at least one seat is required to make a reservation")
)

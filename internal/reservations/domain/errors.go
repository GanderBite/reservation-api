package domain

import "errors"

var (
	ErrReservationStatusAlreadyApplied = errors.New("reservation already has that status applied")
	ErrReservationAlreadyExpired       = errors.New("reservation is already expired")
	ErrReservationNotFound             = errors.New("reservation was not found")
	ErrSeatsAlreadyReserved            = errors.New("one or more selected seats are already reserved")
	ErrMissingSeats                    = errors.New("at least one seat is required to make a reservation")
)

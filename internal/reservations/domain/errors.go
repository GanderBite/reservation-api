package domain

import "errors"

var (
	ErrSeatsAlreadyReserved = errors.New("one or more selected seats are already reserved")
)

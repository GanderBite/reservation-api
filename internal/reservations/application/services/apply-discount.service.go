package services

type ApplyDiscountService interface {
	ApplyDiscountBySeatsCount(seatsCount int) string
}

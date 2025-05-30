package services

type ApplyDiscountService struct{}

func NewApplyDiscountService() *ApplyDiscountService {
	return &ApplyDiscountService{}
}

func (s *ApplyDiscountService) ApplyDiscountBySeatsCount(seatsCount int) string {
	if seatsCount == 2 {
		return "PAIR"
	}

	if seatsCount >= 5 {
		return "BIG_FAM"
	}

	return ""
}

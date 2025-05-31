package services_test

import (
	"testing"

	"github.com/GanderBite/reservation-api/internal/reservations/infrastructure/services"
)

func TestApplyDiscountBySeatsCount(t *testing.T) {
	service := services.NewApplyDiscountService()

	tests := []struct {
		name         string
		seatsCount   int
		expectedCode string
	}{
		{
			name:         "No discount for 1 seat",
			seatsCount:   1,
			expectedCode: "",
		},
		{
			name:         "PAIR discount for 2 seats",
			seatsCount:   2,
			expectedCode: "PAIR",
		},
		{
			name:         "No discount for 3 seats",
			seatsCount:   3,
			expectedCode: "",
		},
		{
			name:         "No discount for 4 seats",
			seatsCount:   4,
			expectedCode: "",
		},
		{
			name:         "BIG_FAM discount for 5 seats",
			seatsCount:   5,
			expectedCode: "BIG_FAM",
		},
		{
			name:         "BIG_FAM discount for 10 seats",
			seatsCount:   10,
			expectedCode: "BIG_FAM",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			discountCode := service.ApplyDiscountBySeatsCount(tt.seatsCount)
			if discountCode != tt.expectedCode {
				t.Errorf("expected %q, got %q", tt.expectedCode, discountCode)
			}
		})
	}
}

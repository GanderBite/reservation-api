package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/GanderBite/reservation-api/internal/pkg/types"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

func TestNewReservationFromDto(t *testing.T) {
	price := types.Price(100)
	res := domain.NewReservationFromDto(price)

	assert.Equal(t, domain.StatusPending, res.Status)
	assert.Equal(t, price, res.Price)
	assert.WithinDuration(t, time.Now(), res.CreatedAt, time.Second)
	assert.WithinDuration(t, res.CreatedAt.Add(time.Hour), res.ExpiresAt, time.Second)
}

func TestApplyDiscount(t *testing.T) {
	price := types.Price(100)
	res := domain.NewReservationFromDto(price)
	discountId := types.Id(uuid.New())
	discountAmount := types.Price(20)

	res.ApplyDiscount(discountId, discountAmount)

	assert.Equal(t, price-discountAmount, res.Price)
	assert.NotNil(t, res.AppliedDiscountCodeId)
	assert.Equal(t, discountId, *res.AppliedDiscountCodeId)
}

func TestUpdateStatus_Success(t *testing.T) {
	res := domain.NewReservationFromDto(100)
	err := res.UpdateStatus(domain.StatusConfirmed)

	assert.NoError(t, err)
	assert.Equal(t, domain.StatusConfirmed, res.Status)
}

func TestUpdateStatus_SameStatus(t *testing.T) {
	res := domain.NewReservationFromDto(100)
	res.Status = domain.StatusConfirmed

	err := res.UpdateStatus(domain.StatusConfirmed)

	assert.ErrorIs(t, err, domain.ErrReservationStatusAlreadyApplied)
}

func TestUpdateStatus_AlreadyExpired(t *testing.T) {
	res := domain.NewReservationFromDto(100)
	res.Status = domain.StatusExpired

	err := res.UpdateStatus(domain.StatusConfirmed)

	assert.ErrorIs(t, err, domain.ErrReservationAlreadyExpired)
}

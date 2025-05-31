package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/GanderBite/reservation-api/internal/reservations/application/dtos"
	usecases "github.com/GanderBite/reservation-api/internal/reservations/application/use-cases"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

type confirmReservationHandler struct {
	updateReservationStatus *usecases.UpdateReservationStatusUseCase
}

func newConfirmReservationHandler(
	updateReservationStatus *usecases.UpdateReservationStatusUseCase,
) *confirmReservationHandler {
	return &confirmReservationHandler{
		updateReservationStatus: updateReservationStatus,
	}
}

// ConfirmReservation godoc
// @Summary Confirm existing reservation
// @Description Confirm given reservation
// @Tags reservations
// @Accept json
// @Produce json
// @Param input body dtos.ConfirmReservationStatusDto true "Reservation Id to confirm"
// @Success 200 {object} response.BoolResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations/confirm [post]
func (h *confirmReservationHandler) Handle(c *gin.Context) {
	var dto dtos.ConfirmReservationStatusDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.updateReservationStatus.Execute(c.Request.Context(), dto.ReservationId, domain.StatusConfirmed)
	if err != nil {
		if errors.Is(err, domain.ErrReservationNotFound) {
			response.Error(c, http.StatusNotFound, err.Error())
		} else if errors.Is(err, domain.ErrReservationAlreadyExpired) {
			response.Error(c, http.StatusBadRequest, err.Error())
		} else if errors.Is(err, domain.ErrReservationStatusAlreadyApplied) {
			response.Error(c, http.StatusConflict, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.Success(c, true)
}

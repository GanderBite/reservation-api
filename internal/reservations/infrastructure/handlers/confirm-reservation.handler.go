package handlers

import (
	"errors"
	"net/http"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/GanderBite/reservation-api/internal/reservations/application/dtos"
	usecases "github.com/GanderBite/reservation-api/internal/reservations/application/use-cases"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
	"github.com/gin-gonic/gin"
)

type confirmReservationHandler struct {
	updateReservationStatus *usecases.UpdateReservationStatusUseCase
}

func newConfirmReservationHandler(updateReservationStatus *usecases.UpdateReservationStatusUseCase) *confirmReservationHandler {
	return &confirmReservationHandler{
		updateReservationStatus: updateReservationStatus,
	}
}

func (h *confirmReservationHandler) Handle(c *gin.Context) {
	var dto dtos.ConfirmReservationStatusDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.updateReservationStatus.Execute(c.Request.Context(), dto.ReservationId, domain.StatusConfirmed)
	if err != nil {
		if errors.Is(domain.ErrReservationNotFound, err) {
			response.Error(c, http.StatusNotFound, err.Error())
		} else if errors.Is(domain.ErrReservationAlreadyExpired, err) {
			response.Error(c, http.StatusBadRequest, err.Error())
		} else if errors.Is(domain.ErrReservationStatusAlreadyApplied, err) {
			response.Error(c, http.StatusConflict, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.Success(c, true)
}

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

type cancelReservationHandler struct {
	updateReservationStatus *usecases.UpdateReservationStatusUseCase
}

func newCancelReservationHandler(updateReservationStatus *usecases.UpdateReservationStatusUseCase) *cancelReservationHandler {
	return &cancelReservationHandler{
		updateReservationStatus: updateReservationStatus,
	}
}

func (h *cancelReservationHandler) Handle(c *gin.Context) {
	var dto dtos.CancelReservationStatusDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.updateReservationStatus.Execute(c.Request.Context(), dto.ReservationId, domain.StatusExpired)
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

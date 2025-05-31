package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
)

// GetReservationDetails godoc
// @Summary Get reservation details
// @Description Returns reservation information and associated seats
// @Tags reservations
// @Produce json
// @Param id path string true "Reservation ID"
// @Success 200 {object} dtos.ReservationSuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations/{id} [get]
func (h *ReservationHandlers) GetReservationDetails(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid id received")

		return
	}

	reservation, err := h.repo.GetReservationDetails(c.Request.Context(), parsedId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	if reservation == nil {
		response.Error(c, http.StatusNotFound, domain.ErrReservationNotFound.Error())
	}

	response.Success(c, reservation)
}

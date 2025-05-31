package handlers

import (
	"net/http"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/GanderBite/reservation-api/internal/reservations/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

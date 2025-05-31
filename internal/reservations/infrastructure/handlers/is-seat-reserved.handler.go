package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	response "github.com/GanderBite/reservation-api/internal/pkg"
)

// IsSeatReserved godoc
// @Summary Checks if given seat is reserved
// @Description Checks if given seat is reserved
// @Tags reservations
// @Produce json
// @Param id path string true "Seat ID"
// @Success 200 {object} response.BoolResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations/is-seat-reserved [get]
func (h *ReservationHandlers) IsSeatReserved(c *gin.Context) {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid id received")

		return
	}

	isReserved, err := h.repo.IsSeatReserved(c.Request.Context(), parsedId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, isReserved)
}

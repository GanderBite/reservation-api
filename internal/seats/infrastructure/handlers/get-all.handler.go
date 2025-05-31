package handlers

import (
	"net/http"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/gin-gonic/gin"
)

// GetAllSeats godoc
// @Summary Gets all seats
// @Description Get all seats
// @Tags seats
// @Produce json
// @Success 200 {object} dtos.GetAllSeatsResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /seats [get]
func (h *SeatHandlers) GetAllSeatsHandler(c *gin.Context) {
	seats, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, seats)
}

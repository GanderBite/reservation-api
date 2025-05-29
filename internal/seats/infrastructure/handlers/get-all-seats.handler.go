package handlers

import (
	"net/http"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/gin-gonic/gin"
)

func (h *SeatHandlers) GetAllSeats(c *gin.Context) {
	seats, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, seats)
}

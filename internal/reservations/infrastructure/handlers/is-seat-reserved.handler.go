package handlers

import (
	"net/http"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

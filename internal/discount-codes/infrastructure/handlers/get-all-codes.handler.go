package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	response "github.com/GanderBite/reservation-api/internal/pkg"
)

// GetAllDiscountCodes godoc
// @Summary Gets all discount codes
// @Description Get all discount codes
// @Tags discount-codes
// @Produce json
// @Success 200 {object} dtos.GetAllDiscountCodesResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /discount-codes [get]
func (h *DiscountCodeHandlers) GetAllCodes(c *gin.Context) {
	discountCodes, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())

		return
	}

	response.Success(c, discountCodes)
}

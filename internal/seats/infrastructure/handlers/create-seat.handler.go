package handlers

import (
	"net/http"

	response "github.com/GanderBite/reservation-api/internal/pkg"
	"github.com/GanderBite/reservation-api/internal/seats/model/dtos"
	usecase "github.com/GanderBite/reservation-api/internal/seats/model/use-cases"
	"github.com/gin-gonic/gin"
)

type createSeatHandler struct {
	createSeatUC *usecase.CreateSeatUseCase
}

func newCreateSeatHandler(createSeatUC *usecase.CreateSeatUseCase) *createSeatHandler {
	return &createSeatHandler{
		createSeatUC: createSeatUC,
	}
}

// CreateSeat godoc
// @Summary Creates a new seat
// @Description Creates a new seat
// @Tags seats
// @Accept json
// @Produce json
// @Param input body dtos.CreateSeatDto true "Seat Input"
// @Success 201 {object} dtos.CreateSeatResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /seats [post]
func (h *createSeatHandler) Handle(c *gin.Context) {
	var dto dtos.CreateSeatDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	seatId, err := h.createSeatUC.Execute(c.Request.Context(), &dto)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, seatId, 201)
}

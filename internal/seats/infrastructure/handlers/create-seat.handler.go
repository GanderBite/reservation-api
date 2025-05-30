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

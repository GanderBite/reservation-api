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

type createReservationHandler struct {
	createReservationUseCase *usecases.CreateReservationUseCase
}

func newCreateReservationHandler(createReservationUC *usecases.CreateReservationUseCase) *createReservationHandler {
	return &createReservationHandler{
		createReservationUseCase: createReservationUC,
	}
}

func (h *createReservationHandler) Handle(c *gin.Context) {
	var dto dtos.CreateReservationDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
	}

	reservationId, err := h.createReservationUseCase.Execute(c.Request.Context(), &dto)
	if err != nil {
		if errors.Is(domain.ErrSeatsAlreadyReserved, err) {
			response.Error(c, http.StatusConflict, err.Error())
		}

		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.Success(c, reservationId, 201)
}

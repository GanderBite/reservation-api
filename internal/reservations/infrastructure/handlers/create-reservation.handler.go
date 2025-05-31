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

// CreateReservation godoc
// @Summary Create a new reservation
// @Description Creates a new reservation with selected seat IDs
// @Tags reservations
// @Accept json
// @Produce json
// @Param input body dtos.CreateReservationDto true "Reservation Input"
// @Success 201 {object} dtos.CreateReservationResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /reservations [post]
func (h *createReservationHandler) Handle(c *gin.Context) {
	var dto dtos.CreateReservationDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	reservationId, err := h.createReservationUseCase.Execute(c.Request.Context(), &dto)
	if err != nil {
		if errors.Is(domain.ErrSeatsAlreadyReserved, err) {
			response.Error(c, http.StatusConflict, err.Error())
		} else if errors.Is(domain.ErrMissingSeats, err) {
			response.Error(c, http.StatusBadRequest, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.Success(c, reservationId, 201)
}

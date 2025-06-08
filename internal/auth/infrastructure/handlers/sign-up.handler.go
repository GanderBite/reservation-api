package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/GanderBite/reservation-api/internal/auth/model/dtos"
	"github.com/GanderBite/reservation-api/internal/auth/model/entities"
	usecases "github.com/GanderBite/reservation-api/internal/auth/model/use-cases"
	response "github.com/GanderBite/reservation-api/internal/pkg"
)

type signUpHandler struct {
	signUpUC *usecases.SignUpUseCase
}

func newSignUpHandler(signUpUC *usecases.SignUpUseCase) *signUpHandler {
	return &signUpHandler{
		signUpUC: signUpUC,
	}
}

// SignUp godoc
// @Summary Creates user and identity
// @Description Creates user and identity
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dtos.SignUpDto true "Sign up input"
// @Success 201 {object} dtos.SignUpResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 409 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/sign-up [post]
func (h *signUpHandler) Handle(c *gin.Context) {
	var dto dtos.SignUpDto

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	jwt, err := h.signUpUC.Execute(c.Request.Context(), &dto)
	if err != nil {
		if errors.Is(err, entities.ErrPasswordNotSecure) {
			response.Error(c, http.StatusBadRequest, err.Error())
		} else if errors.Is(err, entities.ErrUserAlreadyExists) {
			response.Error(c, http.StatusConflict, err.Error())
		} else {
			response.Error(c, http.StatusInternalServerError, err.Error())
		}

		return
	}

	response.Success(c, jwt, 201)
}

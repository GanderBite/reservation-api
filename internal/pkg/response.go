package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"  example:"error"`
	Message string `json:"message" example:"error message"`
}

type BoolResponse struct {
	Status string `json:"status" example:"success"`
	Data   bool   `json:"data"   example:"true"`
}

func Success[T any](c *gin.Context, data T, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	c.JSON(code, SuccessResponse[T]{
		Status: "success",
		Data:   data,
	})
}

func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{
		Status:  "error",
		Message: message,
	})
}

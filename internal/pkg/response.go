package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse[T any] struct {
	Status string `json:"status"` // always "success"
	Data   T      `json:"data"`   // typed payload
}

type ErrorResponse struct {
	Status  string `json:"status"`  // always "error"
	Message string `json:"message"` // simple error message
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

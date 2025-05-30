package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})

		// Discount codes
		v1.GET("/discount-codes", app.discountCodes.Handlers.GetAllCodes)

		// Seats
		v1.POST("/seats", app.seats.Handlers.CreateSeatHandler.Handle)

		// Reservations
		v1.POST("/reservations", app.reservations.Handlers.CreateReservationHandler.Handle)

	}

	return g
}

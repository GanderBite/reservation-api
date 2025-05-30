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

		v1.POST("/reservations", app.reservations.Handlers.CreateReservationHandler.Handle)

		v1.POST("/seats", app.seats.Handlers.CreateSeatHandler.Handle)
	}

	return g
}

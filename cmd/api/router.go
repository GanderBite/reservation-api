package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/GanderBite/reservation-api/docs"
)

func (app *application) routes() http.Handler {
	g := gin.Default()

	v1 := g.Group("/api/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})

		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Discount codes
		v1.GET("/discount-codes", app.discountCodes.Handlers.GetAllCodes)

		// Seats
		v1.GET("/seats", app.seats.Handlers.GetAllSeatsHandler)
		v1.POST("/seats", app.seats.Handlers.CreateSeatHandler.Handle)

		// Reservations
		v1.GET("/reservations/:id", app.reservations.Handlers.GetReservationDetails)
		v1.GET("/reservations/is-seat-reserved/:id", app.reservations.Handlers.IsSeatReserved)
		v1.POST("/reservations", app.reservations.Handlers.CreateReservationHandler.Handle)
		v1.POST("/reservations/confirm", app.reservations.Handlers.ConfirmReservationHandler.Handle)
		v1.POST("/reservations/cancel", app.reservations.Handlers.CancelReservationHandler.Handle)
	}

	return g
}

package routes

import (
	"the-car-wash-directory/internal/handlers"
	"the-car-wash-directory/internal/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AddRoutes maps all routes in one place, so you can see the entire API surface at a glance.
func AddRoutes(
	engine *gin.Engine,
	logger *zap.SugaredLogger,
	carWashService *services.CarWashService,
) {
	engine.GET("/healthz", func(c *gin.Context) {
		c.String(200, "OK")
	})

	engine.GET("/", handlers.ComingSoonHandler)

	// Car Wash Routes
	/* engine.GET("/wash/status/:id", handlers.WashStatusHandler(logger, carWashService))
	engine.POST("/wash/new", handlers.NewWashHandler(logger, carWashService)) */

	// Feedback Form (only needs logger, not carWashService)
	engine.POST("/submit-feedback", handlers.HandleFeedbackForm(logger))

	engine.NoRoute(handlers.NotFoundHandler)
}

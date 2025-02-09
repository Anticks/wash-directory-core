package server

import (
	"the-car-wash-directory/internal/server/routes"
	"the-car-wash-directory/internal/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewServer is the "big constructor" that returns *gin.Engine
// (the main http.Handler for the entire service).
func NewServer(
	logger *zap.SugaredLogger,
	carWashService *services.CarWashService,
) *gin.Engine {
	// Typically you'd use gin.New() for a blank engine,
	// or gin.Default() for built-in logging/recovery middlewares
	engine := gin.New()

	logger.Infow("NewServer created")

	// Global middlewares (top-level)
	engine.Use(gin.Logger())   // or a custom logger middleware
	engine.Use(gin.Recovery()) // handle panics

	// Serve static files (CSS, JS, images, etc.)
	engine.Static("/static", "./static")

	// Setup routes in one place
	routes.AddRoutes(engine, logger, carWashService)

	return engine
}

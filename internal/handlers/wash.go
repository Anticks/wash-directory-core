package handlers

import (
	"context"
	"the-car-wash-directory/internal/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NewWashHandler handles starting a new car wash.
func NewWashHandler(logger *zap.SugaredLogger, carWashService *services.CarWashService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		if err := carWashService.StartWash(ctx); err != nil {
			logger.Errorf("Failed to start wash: %v", err)
			c.JSON(500, gin.H{"error": "Failed to start wash"})
			return
		}

		c.JSON(200, gin.H{"message": "Car wash started successfully"})
	}
}

// WashStatusHandler handles retrieving the status of a car wash.
func WashStatusHandler(logger *zap.SugaredLogger, carWashService *services.CarWashService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		washID := c.Param("id")

		status, err := carWashService.GetWashStatus(ctx, washID)
		if err != nil {
			logger.Errorf("Failed to get wash status: %v", err)
			c.JSON(500, gin.H{"error": "Failed to get wash status"})
			return
		}

		c.JSON(200, gin.H{"status": status})
	}
}

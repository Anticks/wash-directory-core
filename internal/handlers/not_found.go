package handlers

import (
	"net/http"
	"the-car-wash-directory/internal/templates"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NotFoundHandler handles 404 Not Found errors
func NotFoundHandler(c *gin.Context) {
	// Log the 404 request for debugging purposes
	zap.L().Warn("404 Not Found", zap.String("path", c.Request.URL.Path))

	// Serve a friendly 404 page
	c.Status(http.StatusNotFound)
	templates.NotFoundPage().Render(c.Request.Context(), c.Writer)
}

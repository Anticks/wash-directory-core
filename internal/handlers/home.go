package handlers

import (
	"the-car-wash-directory/internal/templates"

	"github.com/gin-gonic/gin"
)

// HomeHandler handles requests to the home page.
func HomeHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	templates.Home("Car Wash").Render(c.Request.Context(), c.Writer)
}

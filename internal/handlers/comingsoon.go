package handlers

import (
	"the-car-wash-directory/internal/templates"

	"github.com/gin-gonic/gin"
)

func ComingSoonHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")

	templates.ComingSoon().Render(c.Request.Context(), c.Writer)
}

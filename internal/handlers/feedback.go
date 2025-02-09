package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	datastar "github.com/starfederation/datastar/sdk/go"
	"go.uber.org/zap"
)

// FeedbackPayload defines the expected JSON structure for feedback.
type FeedbackPayload struct {
	Email   string `json:"email"`
	Message string `json:"message"`
	Toast   string `json:"toast"`
}

// HandleFeedbackForm returns a handler function for processing the feedback form.
func HandleFeedbackForm(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload FeedbackPayload
		update := gabs.New()

		// Bind the JSON payload into our struct.
		if err := c.ShouldBindJSON(&payload); err != nil {
			logger.Warnf("Failed to bind JSON: %v", err)
			update.Set("Invalid input.", "response")
			c.Status(http.StatusBadRequest)
			datastar.NewSSE(c.Writer, c.Request).MarshalAndMergeSignals(update)
			return
		}

		logger.Infof("Received feedback with email: %s and message: %s", payload.Email, payload.Message)

		// Validate input.
		if payload.Email == "" || payload.Message == "" {
			logger.Warn("Missing email or message in feedback form")
			// update.Set("Email and Message are required.", "response")
			c.Status(http.StatusBadRequest)
			datastar.NewSSE(c.Writer, c.Request).MarshalAndMergeSignals(update)
			return
		}

		// Configure SendGrid email.
		from := mail.NewEmail("Car Wash Directory", "hello@washdirectory.com")
		subject := "New Feedback Submission"
		to := mail.NewEmail("Admin", "hello@washdirectory.com")
		plainTextContent := fmt.Sprintf("From: %s\nMessage:\n%s", payload.Email, payload.Message)
		htmlContent := fmt.Sprintf("<p><strong>From:</strong> %s</p><p><strong>Message:</strong><br>%s</p>", payload.Email, payload.Message)

		messageToSend := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

		response, err := client.Send(messageToSend)
		if err != nil {
			logger.Errorf("Failed to send email: %v", err)
			// update.Set("Failed to send email. Please try again.", "response")
			c.Status(http.StatusInternalServerError)
			datastar.NewSSE(c.Writer, c.Request).MarshalAndMergeSignals(update)
			return
		}

		if response.StatusCode >= 200 && response.StatusCode < 300 {
			logger.Infof("Feedback email sent successfully from %s", payload.Email)
			// update.Set("Thank you for your feedback! We'll get back to you soon.", "response")
			// Update the toast signal with a plain text message.
			update.Set("Message sent successfully.", "toast")
			c.Status(http.StatusOK)
		} else {
			logger.Errorf("SendGrid error: %d - %s", response.StatusCode, response.Body)
			update.Set("Something went wrong. Please try again.", "response")
			c.Status(http.StatusInternalServerError)
		}

		datastar.NewSSE(c.Writer, c.Request).MarshalAndMergeSignals(update)
	}
}

package middleware

import (
	"log"
	"net/http"

	"github.com/Tretorhate/university-management-system/pkg/errors"
	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware that handles errors and returns consistent responses
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if there are any errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			
			// Check if it's an AppError wow
			if appErr, ok := errors.IsAppError(err); ok {
				c.JSON(appErr.Code, gin.H{
					"error": appErr.Message,
				})
				return
			}

			// Log the error
			log.Printf("Unhandled error: %v", err)

			// Return a generic error response
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "An unexpected error occurred",
			})
		}
	}
} 
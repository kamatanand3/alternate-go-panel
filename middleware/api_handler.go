package middleware

import (
	"altpanel/helpers"
	"github.com/gin-gonic/gin"
	"time"
)

func ApiHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = "N/A"
		}

		sourceApp := c.GetHeader("X-Source-App")

		helpers.SetRequestId(requestID)
		helpers.SetSourceApp(sourceApp)
		helpers.SetChannel("api") // channel set here

		helpers.Info("ApiHandler", "Request", "request_url", c.Request.RequestURI)

		c.Next()

		duration := time.Since(start)
		helpers.Info("ApiHandler", "Response", "execution_time", duration.String())
	}
}

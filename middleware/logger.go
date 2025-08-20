package middleware

import (
	"bytes"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// CustomWriter captures response body
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b) // copy response body
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	// Open file in append mode
	file, err := os.OpenFile("logs/apilog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Could not open log file:", err)
	}
	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		// Capture request
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = "N/A"
		}

		// Copy request body
		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			requestBody = bodyBytes
			// Reset body so Gin can read it again
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Capture response
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		start := time.Now()
		c.Next() // process request
		latency := time.Since(start)

		// Log to file
		logger.Printf("| %s | %s %s | RequestID=%s | ReqBody=%s | Status=%d | Resp=%s | Time=%v",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			requestID,
			string(requestBody),
			c.Writer.Status(),
			blw.body.String(),
			latency,
		)
	}
}

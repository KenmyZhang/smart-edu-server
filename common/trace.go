package common

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"smart-edu-server/log"
)

func Logger(notLogged ...string) gin.HandlerFunc {
	var skip map[string]struct{}

	if length := len(notLogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notLogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			log.Info(fmt.Sprintf("%3d | %13v | %-15s | %-7s %s %s",
				statusCode,
				latency,
				clientIP,
				method,
				path,
				comment,
			))
		}
	}
}

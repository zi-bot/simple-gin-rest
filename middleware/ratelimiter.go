package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitter() gin.HandlerFunc {
	// 100 per minute
	capacity := 100
	duration := time.Minute
	bucket := ratelimit.NewBucket(duration, int64(capacity))
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}

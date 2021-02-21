package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// 全局令牌桶限流
func RateLimitByTokenBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		ratelimit.NewBucket(time.Second,0)
		c.Next()
	}
}

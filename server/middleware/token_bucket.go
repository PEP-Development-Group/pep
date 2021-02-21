package middleware

import (
	"github.com/gin-gonic/gin"
)

// 全局令牌桶限流
func RateLimitByTokenBucket() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

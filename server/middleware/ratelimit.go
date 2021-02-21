package middleware

import (
	"github.com/gin-gonic/gin"
)

// 限制单IP每秒请求速率
func RateLimitByIP() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

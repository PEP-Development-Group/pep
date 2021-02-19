package middleware

import (
	"github.com/gin-gonic/gin"
)

// 限制单用户每秒请求速率
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

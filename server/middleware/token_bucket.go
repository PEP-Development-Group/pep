package middleware

import (
	"github.com/gin-gonic/gin"
	"pep/global"
)

// 全局令牌桶限流
func RateLimitByTokenBucket() gin.HandlerFunc {
	return func(c *gin.Context) {
		if remain := global.GVA_BUCKET.TakeAvailable(1); remain == 0 {
			c.AbortWithStatus(403)
		}
		c.Next()
	}
}
